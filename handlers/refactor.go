package handlers

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/minhajuddinkhan/gatewaygo/queue"

	"github.com/sirupsen/logrus"

	"github.com/minhajuddinkhan/todogo/utils"
	nsq "github.com/nsqio/go-nsq"

	"github.com/darahayes/go-boom"

	"github.com/minhajuddinkhan/gatewaygo/constants"
	"github.com/minhajuddinkhan/gatewaygo/targets"

	"github.com/jinzhu/gorm"
)

type Feeds struct {
	DataModel  string
	Event      string
	EndpointID uint
}

//RefactoredHandler RefactoredHandler
func RefactoredHandler(db *gorm.DB, producer *nsq.Producer) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		b, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		ctx, cancelFunc := context.WithTimeout(r.Context(), 1*time.Second)
		defer cancelFunc()
		feeds := []Feeds{
			{
				DataModel:  "patient",
				Event:      "New",
				EndpointID: 1,
			},
			{
				DataModel:  "practitioner",
				Event:      "New",
				EndpointID: 2,
			},
			{
				DataModel:  "appointment",
				Event:      "New",
				EndpointID: 3,
			},
			{
				DataModel:  "encounter",
				Event:      "New",
				EndpointID: 4,
			},
		}

		finished := make(chan bool, 1)
		errChannel := make(chan error)

		var wg sync.WaitGroup
		wg.Add(len(feeds))

		EndpointIds := []uint{1, 2, 4, 3}
		nsqMessage := queue.NSQMessage{
			EndpointIDs: EndpointIds,
			Fragments:   []queue.Fragment{},
		}
		for _, f := range feeds {
			go func(f Feeds) {
				defer wg.Done()
				t := targets.NewDefaultTarget(f.DataModel, f.Event)
				res, err := t.ToFHIR(b)
				if err != nil {
					errChannel <- err
					return
				}
				nsqMessage.Fragments = append(nsqMessage.Fragments, queue.Fragment{
					DataModel:  f.DataModel,
					EndpointID: f.EndpointID,
					Data:       res,
				})

			}(f)
		}

		go func() { wg.Wait(); close(finished) }()

		select {
		case <-ctx.Done():
			utils.Respond(w, "Timed out!")
			return
		case <-finished:
		case err := <-errChannel:
			boom.BadRequest(w, err)
			return

		}

		b, _ = json.Marshal(nsqMessage)
		err := producer.Publish(constants.TOPIC, b)
		if err != nil {
			logrus.Error("cudnt publish", err.Error())
		} else {
			logrus.Info("MSG PUBLISHED!")
		}

		utils.Respond(w, nsqMessage)

	}
}

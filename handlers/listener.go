package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/minhajuddinkhan/gatewaygo/constants"
	"github.com/minhajuddinkhan/gatewaygo/queue"
	"github.com/minhajuddinkhan/gatewaygo/targets"
	nsq "github.com/nsqio/go-nsq"

	"github.com/darahayes/go-boom"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/minhajuddinkhan/gatewaygo/auth"
	"github.com/minhajuddinkhan/gatewaygo/mappers"
	"github.com/minhajuddinkhan/gatewaygo/models"

	"github.com/minhajuddinkhan/todogo/utils"
)

type RedoxRequestMeta struct {
	Meta struct {
		DataModel string
		EventType string
		Source    struct {
			ID string
		}
	}
}

type Payload struct {
	Source    uint `json:"source"`
	DataModel string
	Event     string
}

//ListenerHandler ListenerHandler
func ListenerHandler(db *gorm.DB, producer *nsq.Producer) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		ctx, cancelFunc := context.WithTimeout(r.Context(), 1*time.Second)
		defer cancelFunc()

		logrus.Info(r.Host)

		err := auth.VerifyToken(db, r.Header.Get("verification-token"))
		if err != nil {
			utils.Respond(w, err.Error())
		}

		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			panic(err)
		}

		var x RedoxRequestMeta

		what := json.Unmarshal(b, &x)
		if what != nil {
			panic(err)
		}

		var redoxSource models.RedoxSources
		if db.Where(`"redoxId" = ?`, x.Meta.Source.ID).Find(&redoxSource).RowsAffected == 0 {
			boom.NotFound(w, "Cannot find redox source")
			return
		}

		payload := Payload{
			Source:    redoxSource.TargetID,
			DataModel: mappers.DataModelMapper[x.Meta.DataModel],
		}
		payload.Event = mappers.EventMapper[payload.DataModel][x.Meta.EventType]

		subscription := models.Subscriptions{
			SourceID: payload.Source,
			Event: models.Events{
				Name: payload.Event,
				DataModel: models.DataModels{
					Name: payload.DataModel,
				},
			},
		}

		err = db.Preload("Source").
			Preload("Target").
			Preload("Event.DataModel").
			Preload("TargetDestination").
			Preload("Endpoint").
			Joins(`left join events e on "eventId" = e.id`).
			Joins(`left join "dataModels" d on ("dataModelId" = d.id)`).
			Where(`"sourceId" = ? and e."name" = ? and d."name" = ?`, subscription.SourceID, subscription.Event.Name, subscription.Event.DataModel.Name).
			Find(&subscription).Error
		if err != nil {
			boom.NotFound(w, err)
			return
		}

		subscription.Endpoint.GetPopulatedEndpoints(db)

		finished := make(chan bool, 1)
		errChannel := make(chan error)
		nsqFragmentCh := make(chan queue.Fragment)

		var wg sync.WaitGroup
		wg.Add(len(subscription.Endpoint.DependentURLs))

		nsqMessage := queue.NSQMessage{
			EndpointIDs: subscription.Endpoint.DependentURLIDs,
			Fragments:   []queue.Fragment{},
		}

		for _, dependentURL := range subscription.Endpoint.DependentURLs {

			var target targets.Target
			if fn, ok := targets.TargetsMap[subscription.Source.Name]; ok {
				target = fn(dependentURL.Event.DataModel.Name, dependentURL.Event.Name)
			} else {

				df := &targets.DefaultTarget{}
				df.New(dependentURL.Event.DataModel.Name, dependentURL.Event.Name)
				target = df
			}

			go func(dependentURL models.Endpoints, target targets.Target, b []byte) {
				defer wg.Done()
				res, err := target.ToFHIR(b)
				if err != nil {
					errChannel <- err
					return
				}
				nsqFragmentCh <- queue.Fragment{
					DataModel:  dependentURL.Event.DataModel.Name,
					EndpointID: dependentURL.ID,
					Data:       res,
					Endpoint:   dependentURL,
					TargetKey:  subscription.Target.Key,
				}

			}(dependentURL, target, b)
		}
		go func() { wg.Wait(); close(finished) }()

		wait := true
		for wait {
			select {
			case f := <-nsqFragmentCh:
				nsqMessage.Fragments = append(nsqMessage.Fragments, f)
			case <-ctx.Done():
				utils.Respond(w, "Timed out!")
				wait = false
				return
			case <-finished:
				fmt.Println("DONE")
				wait = false
			case err := <-errChannel:
				boom.BadRequest(w, err)
				return

			}

		}

		orderedFragments := []queue.Fragment{}
		for _, endpointID := range nsqMessage.EndpointIDs {
			for _, nestedF := range nsqMessage.Fragments {
				if endpointID == nestedF.EndpointID {
					orderedFragments = append(orderedFragments, nestedF)
				}
			}
		}
		nsqMessage.Fragments = orderedFragments
		b, _ = json.Marshal(nsqMessage)
		err = producer.Publish(constants.TOPIC, b)
		if err != nil {
			logrus.Error("cudnt publish", err.Error())
		} else {
			logrus.Info("MSG PUBLISHED!")
		}

		utils.Respond(w, struct {
			Done bool
		}{
			true,
		})

	}
}

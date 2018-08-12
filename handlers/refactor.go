package handlers

import (
	"context"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/minhajuddinkhan/todogo/utils"

	"github.com/darahayes/go-boom"

	"github.com/minhajuddinkhan/gatewaygo/targets"

	"github.com/jinzhu/gorm"
)

type Feeds struct {
	DataModel string
	Event     string
}

//RefactoredHandler RefactoredHandler
func RefactoredHandler(db *gorm.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		b, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		ctx, cancelFunc := context.WithTimeout(r.Context(), 1*time.Second)
		defer cancelFunc()
		feeds := []Feeds{
			{
				DataModel: "patient",
				Event:     "New",
			},
			{
				DataModel: "practitioner",
				Event:     "New",
			},
			{
				DataModel: "appointment",
				Event:     "New",
			},
			{
				DataModel: "encounter",
				Event:     "New",
			},
		}

		var fhirResults []interface{}
		finished := make(chan bool, 1)
		errChannel := make(chan error)

		var wg sync.WaitGroup
		wg.Add(len(feeds))

		for _, f := range feeds {
			go func(f Feeds) {
				defer wg.Done()
				t := targets.NewDefaultTarget(f.DataModel, f.Event)
				res, err := t.ToFHIR(b)
				if err != nil {
					errChannel <- err
					return
				}
				fhirResults = append(fhirResults, res)

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

		utils.Respond(w, fhirResults)

	}
}

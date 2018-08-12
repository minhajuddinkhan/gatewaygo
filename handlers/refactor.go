package handlers

import (
	"io/ioutil"
	"net/http"
	"sync"

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
				Event:     "New1",
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
		case <-finished:
		case err := <-errChannel:
			boom.BadRequest(w, err)
			return

		}

		utils.Respond(w, fhirResults)

	}
}

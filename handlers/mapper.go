package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/minhajuddinkhan/gatewaygo/fhir"

	"github.com/minhajuddinkhan/todogo/utils"

	"github.com/minhajuddinkhan/gatewaygo/redox/models/common"
	"github.com/minhajuddinkhan/gatewaygo/redox/models/scheduling"

	"github.com/jinzhu/gorm"
)

func MapperHandler(db *gorm.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var wg sync.WaitGroup
		b, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		var y struct {
			Meta common.Meta
		}
		json.Unmarshal(b, &y)

		var patient, practitioner, appointment, encounter interface{}
		wg.Add(4)
		if y.Meta.DataModel == "Scheduling" && y.Meta.EventType == "New" {

			RedoxScheduling := scheduling.New{}
			err := json.Unmarshal(b, &RedoxScheduling)
			if err != nil {
				panic("SOMETHING WENT WRONG" + err.Error())
			}

			go func() {
				defer wg.Done()
				patient, err = fhir.NewFHIRPatient(RedoxScheduling)
				if err != nil {
					panic(err)
				}
			}()

			go func() {
				defer wg.Done()
				practitioner, err = fhir.NewFHIRPractitioner(RedoxScheduling)
				if err != nil {
					panic(err)
				}

			}()

			go func() {
				defer wg.Done()
				appointment, err = fhir.NewAppointment(RedoxScheduling)
				if err != nil {
					panic(err)
				}

			}()

			go func() {
				defer wg.Done()
				encounter, err = fhir.NewFHIREncounter(RedoxScheduling)
				if err != nil {
					panic(err)
				}
			}()

			wg.Wait()

			utils.Respond(w, struct {
				Patient     interface{}
				Provider    interface{}
				Appointment interface{}
				Encounter   interface{}
			}{
				patient,
				practitioner,
				appointment,
				encounter,
			})
			return

		}

	}
}

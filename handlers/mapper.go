package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/minhajuddinkhan/fhir/models"
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

		var fhirPatient models.Patient
		var fhirPractitioner models.Practitioner
		var fhirAppointment models.Appointment

		wg.Add(3)
		if y.Meta.DataModel == "Scheduling" && y.Meta.EventType == "New" {

			RedoxScheduling := scheduling.New{}
			err := json.Unmarshal(b, &RedoxScheduling)
			if err != nil {
				panic("SOMETHING WENT WRONG" + err.Error())
			}

			go func() {
				defer wg.Done()
				fhirPatient = fhir.NewFHIRPatient(RedoxScheduling)
			}()

			go func() {
				defer wg.Done()
				fhirPractitioner = fhir.NewFHIRPractitioner(RedoxScheduling)
			}()

			go func() {
				defer wg.Done()
				fhirAppointment = fhir.NewAppointment(RedoxScheduling)
			}()

			wg.Wait()

			app, _ := fhirAppointment.GetBSON()

			utils.Respond(w, struct {
				Patient     models.Patient
				Provider    models.Practitioner
				Appointment interface{}
			}{
				fhirPatient,
				fhirPractitioner,
				app,
			})
			return

		}

	}
}

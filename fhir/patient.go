package fhir

import (
	"encoding/json"
	"time"

	"github.com/minhajuddinkhan/fhir/models"
	"github.com/minhajuddinkhan/gatewaygo/redox/models/scheduling"
)

//NewFHIRPatient NewFHIRPatient
func NewFHIRPatient(b []byte) ([]byte, error) {

	var RedoxScheduling scheduling.New

	json.Unmarshal(b, &RedoxScheduling)

	fhirDateLayout := "2006-01-02"
	p := models.Patient{
		Identifier: []models.Identifier{
			{
				Value:  RedoxScheduling.Patient.Identifiers[0].ID,
				System: RedoxScheduling.Patient.Identifiers[0].IDType,
			},
		},
		Name: []models.HumanName{
			{
				Given:  []string{RedoxScheduling.Patient.Demographics.FirstName + " " + RedoxScheduling.Patient.Demographics.MiddleName},
				Family: []string{RedoxScheduling.Patient.Demographics.LastName},
			},
		},
		Address: []models.Address{
			{
				City:    RedoxScheduling.Patient.Demographics.Address.City,
				Country: RedoxScheduling.Patient.Demographics.Address.Country,
				Line: []string{
					RedoxScheduling.Patient.Demographics.Address.StreetAddress,
				},
				PostalCode: RedoxScheduling.Patient.Demographics.Address.ZIP,
				State:      RedoxScheduling.Patient.Demographics.Address.State,
			},
		},
		DeceasedBoolean: RedoxScheduling.Patient.Demographics.IsDeceased,
		BirthDate: &models.FHIRDateTime{
			Time: func() time.Time {
				t, _ := time.Parse(fhirDateLayout, RedoxScheduling.Patient.Demographics.DOB)
				return t
			}(),
		},
		Gender: RedoxScheduling.Patient.Demographics.Sex,
		Telecom: []models.ContactPoint{
			{
				System: "phone",
				Value:  RedoxScheduling.Patient.Demographics.PhoneNumber.Home,
				Use:    "home",
			},
			{
				System: "phone",
				Value:  RedoxScheduling.Patient.Demographics.PhoneNumber.Mobile,
				Use:    "mobile",
			},
			{
				System: "phone",
				Value:  RedoxScheduling.Patient.Demographics.PhoneNumber.Office,
				Use:    "office",
			},
		},
		Communication: []models.PatientCommunicationComponent{
			{
				Language: &models.CodeableConcept{
					Text: RedoxScheduling.Patient.Demographics.Language,
				},
			},
		},
	}

	return p.MarshalJSON()

}

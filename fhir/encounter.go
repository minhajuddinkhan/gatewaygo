package fhir

import (
	"encoding/json"
	"strings"

	"github.com/minhajuddinkhan/fhir/models"
	"github.com/minhajuddinkhan/gatewaygo/redox/models/scheduling"
)

//NewFHIREncounter NewFHIREncounter
func NewFHIREncounter(bytes []byte) ([]byte, error) {

	var redoxPayload scheduling.New
	json.Unmarshal(bytes, &redoxPayload)

	e := models.Encounter{
		Identifier: []models.Identifier{
			{
				Value: redoxPayload.Visit.VisitNumber,
			},
		},
		Appointment: &models.Reference{
			Reference: redoxPayload.Visit.VisitNumber,
			Display:   "Appointment",
		},
		Status: strings.ToLower(redoxPayload.Visit.Status),
		Length: &models.Quantity{
			Value: func() *float64 {
				f := float64(*redoxPayload.Visit.Duration)
				return &f
			}(),
		},
		Participant: []models.EncounterParticipantComponent{
			{
				Individual: &models.Reference{
					Reference: redoxPayload.Visit.AttendingProvider.ID,
					Display:   "Practitioner",
				},
			},
			{
				Individual: &models.Reference{
					Reference: redoxPayload.Patient.Identifiers[0].ID,
					Display:   "Patient",
				},
			},
		},
		Location: []models.EncounterLocationComponent{
			{
				Location: &models.Reference{
					Display: redoxPayload.Visit.Location.Facility,
				},
			},
		},
		Reason: []models.CodeableConcept{
			{
				Text: redoxPayload.Visit.Reason,
			},
		},
	}

	return e.MarshalJSON()

}

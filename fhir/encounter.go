package fhir

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/minhajuddinkhan/fhir/models"
	"github.com/minhajuddinkhan/gatewaygo/redox/models/scheduling"
)

//NewFHIREncounter NewFHIREncounter
func NewFHIREncounter(bytes []byte) ([]byte, error) {

	var redoxPayload scheduling.New
	err := json.Unmarshal(bytes, &redoxPayload)
	if err != nil {
		return []byte{}, err
	}

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
				var itg float64
				u64, _ := strconv.ParseUint(redoxPayload.Visit.Duration, 10, 16)
				itg = float64(u64)
				return &itg
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

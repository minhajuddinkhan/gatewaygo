package targets

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/minhajuddinkhan/gatewaygo/fhir"
	"github.com/minhajuddinkhan/gatewaygo/queue"
)

func getMapOfFunc(mapString string, mapFunc func(b []byte, destinationCode string) ([]byte, error)) map[string]func(b []byte, destinationCode string) ([]byte, error) {

	m := make(map[string]func([]byte, string) ([]byte, error))
	m[mapString] = mapFunc
	return m
}

var (
	//DefaultMapper DefaultMapper
	DefaultMapper = map[string]map[string]func(b []byte, destinationCode string) ([]byte, error){

		"patient":      getMapOfFunc("New", fhir.NewFHIRPatient),
		"practitioner": getMapOfFunc("New", fhir.NewFHIRPractitioner),
		"appointment":  getMapOfFunc("Schedule", fhir.NewAppointment),
		"encounter":    getMapOfFunc("New", fhir.NewFHIREncounter),
	}
)

//DefaultTarget DefaultTarget
type DefaultTarget struct {
	name       string
	DataModel  string
	Event      string
	AuthParams string
}

//ToFHIR ToFHIR
func (d *DefaultTarget) ToFHIR(b []byte, destinationCode string) ([]byte, error) {
	var fhir []byte
	if fn, ok := DefaultMapper[d.DataModel][d.Event]; ok {
		result, err := fn(b, destinationCode)
		if err != nil {
			return fhir, fmt.Errorf("Cannot Map for Default Target. Error: %s", err.Error())
		}
		return result, nil
	}
	return fhir, fmt.Errorf("Mapper not configured for Default Mapper {DataModel: %s, Event: %s}", d.DataModel, d.Event)

}

//GetAttribute GetAttribute
func (d *DefaultTarget) GetAttribute(key string) (string, error) {

	switch key {
	case "AuthParams":
		return d.AuthParams, nil
	default:
		return "", errors.New("key not found")
	}
}

//Execute Execute
func (d *DefaultTarget) Execute(payload *queue.NSQMessage) {

}

//New New
func (d *DefaultTarget) New(dataModel, event, authParams string) *DefaultTarget {

	if len(authParams) != 0 {
		err := json.Unmarshal([]byte(authParams), &d.AuthParams)
		if err != nil {
			panic(err.Error())
		}
	}
	d.name = "default"
	d.DataModel = dataModel
	d.Event = event
	return d

}

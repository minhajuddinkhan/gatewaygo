package targets

import (
	"encoding/json"
	"fmt"

	"github.com/minhajuddinkhan/gatewaygo/queue"

	"github.com/minhajuddinkhan/gatewaygo/fhir"
)

var (
	//DefaultMapper DefaultMapper
	DefaultMapper = map[string]map[string]func(b []byte) ([]byte, error){

		"appointment": map[string]func(b []byte) ([]byte, error){
			"Schedule": func(b []byte) ([]byte, error) {
				return fhir.NewAppointment(b)
			},
		},
		"patient": map[string]func(b []byte) ([]byte, error){
			"New": func(b []byte) ([]byte, error) {
				return fhir.NewFHIRPatient(b)
			},
		},
		"practitioner": map[string]func(b []byte) ([]byte, error){
			"New": func(b []byte) ([]byte, error) {
				return fhir.NewFHIRPractitioner(b)
			},
		},
		"encounter": map[string]func(b []byte) ([]byte, error){
			"New": func(b []byte) ([]byte, error) {
				return fhir.NewFHIREncounter(b)
			},
		},
	}
)

//DefaultTarget DefaultTarget
type DefaultTarget struct {
	name      string
	DataModel string
	Event     string
}

//New New
func (d *DefaultTarget) New(dataModel, event string) *DefaultTarget {

	d.name = "default"
	d.DataModel = dataModel
	d.Event = event
	return d

}

//ToFHIR ToFHIR
func (d *DefaultTarget) ToFHIR(b []byte) ([]byte, error) {
	var fhir []byte
	if fn, ok := DefaultMapper[d.DataModel][d.Event]; ok {
		result, err := fn(b)
		if err != nil {
			return fhir, fmt.Errorf("Cannot Map for Default Target. Error: %s", err.Error())
		}
		return result, nil
	}
	return fhir, fmt.Errorf("Mapper not configured for Default Mapper {DataModel: %s, Event: %s}", d.DataModel, d.Event)

}

//Execute Execute
func (d *DefaultTarget) Execute(payload *queue.NSQMessage) {

	for _, o := range payload.Fragments {
		fmt.Println(o.Endpoint.Params)
		var ep struct {
			Method string
		}
		json.Unmarshal([]byte(o.Endpoint.Params), &ep)
		fmt.Println(ep.Method)

		// if ep.Method == "GET" {
		// 	method = "GET"
		// } else {
		// 	method = "POST"
		// }

		// url = o.Endpoint.URL

	}

}

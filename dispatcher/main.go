package main

import (
	"encoding/json"
	"fmt"

	"github.com/minhajuddinkhan/fhir/models"
	"github.com/minhajuddinkhan/gatewaygo/queue"

	"github.com/gorilla/mux"
	"github.com/minhajuddinkhan/gatewaygo/constants"
	"github.com/minhajuddinkhan/todogo/server"
	"github.com/nsqio/go-nsq"
)

func main() {

	fmt.Println("WHOO!")

	r := mux.NewRouter()

	consumer, err := nsq.NewConsumer(constants.TOPIC, constants.CHANNEL, nsq.NewConfig())
	if err != nil {
		panic(err)
	}

	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {

		// stringToDateTimeHook := func(
		// 	f reflect.Type,
		// 	t reflect.Type,
		// 	data interface{}) (interface{}, error) {
		// 	fmt.Println("INTERFAC", data)
		// 	if t == reflect.TypeOf(time.Time{}) && f == reflect.TypeOf("") {
		// 		return time.Parse(time.RFC3339, data.(string))
		// 	}

		// 	return data, nil
		// }
		// config := mapstructure.DecoderConfig{
		// 	DecodeHook: stringToDateTimeHook,
		// 	Result:     &patient,
		// }

		// decoder, err := mapstructure.NewDecoder(&config)
		// if err != nil {
		// 	fmt.Println("ERR BRO", err.Error())
		// }

		patient := models.Patient{}

		nsqMessage := queue.NSQMessage{}
		json.Unmarshal(message.Body, &nsqMessage)

		for _, m := range nsqMessage.Fragments {
			if m.DataModel == "patient" {
				json.Unmarshal(m.Data, &patient)
			}
		}

		fmt.Println(patient.BirthDate)

		// for k, v := range p {
		// 	switch val := v.(type) {
		// 	case string:
		// 		fmt.Println("GO!", val, k, v)
		// 	case nil:
		// 		fmt.Println("ITS NIL")
		// 	default:
		// 		fmt.Println("CANT IDENTIFy")
		// 	}

		// }
		// for k, v := range nsqMessage.Fragments[0].Data {
		// 	fmt.Println("Key", k)
		// 	fmt.Println("VALUE", v)
		// }

		// orderedFragments := []queue.Fragment{}
		// for _, endpointID := range nsqMessage.EndpointIDs {
		// 	for _, nestedF := range nsqMessage.Fragments {
		// 		if endpointID == nestedF.EndpointID {
		// 			orderedFragments = append(orderedFragments, nestedF)
		// 		}
		// 	}

		// }
		// appointment := models.Appointment{}
		// for _, x := range orderedFragments {
		// 	if x.DataModel == "appointment" {
		// 		fmt.Println(x.Data)
		// 		bytes, err := json.Marshal(x.Data)
		// 		if err != nil {
		// 			fmt.Println("COUDNLT MARSHAL", err.Error())
		// 		}
		// 		err = json.Unmarshal(bytes, &appointment)
		// 		if err != nil {
		// 			fmt.Println("COULDNt UNMARSHAL", appointment)
		// 		}

		// 	}
		// }

		// fmt.Println(appointment)

		message.Finish()
		return nil
	}))

	err = consumer.ConnectToNSQD(":4150")
	if err != nil {
		panic(err)
	}

	svr := server.NewServer()
	svr.Listen(":8080", r)

}

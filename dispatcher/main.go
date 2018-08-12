package main

import (
	"encoding/json"
	"fmt"

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

		nsqMessage := queue.NSQMessage{}
		json.Unmarshal(message.Body, &nsqMessage)

		orderedFragments := []queue.Fragment{}
		for _, endpointID := range nsqMessage.EndpointIDs {
			for _, nestedF := range nsqMessage.Fragments {
				if endpointID == nestedF.EndpointID {
					orderedFragments = append(orderedFragments, nestedF)
				}
			}
		}

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

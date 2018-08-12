package main

import (
	"fmt"
	"log"

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

		log.Printf("Got a message: %v", message)
		return nil
	}))

	err = consumer.ConnectToNSQD(":4150")
	if err != nil {
		panic(err)
	}

	svr := server.NewServer()
	svr.Listen(":8080", r)

}

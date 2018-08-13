package main

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/minhajuddinkhan/gatewaygo/queue"

	"github.com/gorilla/mux"
	"github.com/minhajuddinkhan/gatewaygo/constants"
	"github.com/minhajuddinkhan/todogo/server"
	"github.com/nsqio/go-nsq"
)

type PostgresJson struct {
	json.RawMessage
}

func (j PostgresJson) Value() (driver.Value, error) {
	return j.MarshalJSON()
}

func (j *PostgresJson) Scan(src interface{}) error {
	if bytes, ok := src.([]byte); ok {
		return json.Unmarshal(bytes, j)

	}
	return errors.New(fmt.Sprint("Failed to unmarshal JSON from DB", src))
}

type EndpointParam struct {
	Method string
}

func main() {

	fmt.Println("WHOO!")

	r := mux.NewRouter()

	consumer, err := nsq.NewConsumer(constants.TOPIC, constants.CHANNEL, nsq.NewConfig())
	if err != nil {
		panic(err)
	}

	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {

		fmt.Println("MESSAGE RECIEVED!")
		nsqMessage := queue.NSQMessage{}
		err := json.Unmarshal(message.Body, &nsqMessage)
		if err != nil {
			fmt.Println(err.Error())
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

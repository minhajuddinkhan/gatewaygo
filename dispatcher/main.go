package main

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	"github.com/minhajuddinkhan/gatewaygo/models"
	"github.com/minhajuddinkhan/gatewaygo/queue"
	"github.com/minhajuddinkhan/gatewaygo/targets"

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
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=dbuser password=dbuser sslmode=disable dbname=dbuser")
	//	db = db.LogMode(true)
	if err != nil {
		panic(err)
	}

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

		var t models.Targets
		if db.Where("key = ?", nsqMessage.Target.Key).Find(&t).RowsAffected == 0 {
			fmt.Println("NO TARGET FOUND")
		}
		var target targets.Target
		if _, ok := targets.TargetsMap[nsqMessage.Target.Key]; !ok {
			logrus.Error("Don't know which target to execute.")
			message.Finish()
			return nil
		}

		target = targets.TargetsMap[nsqMessage.Target.Key]("", "", t.AuthParams)
		target.Execute(&nsqMessage)

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

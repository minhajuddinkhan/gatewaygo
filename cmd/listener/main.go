package main

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/minhajuddinkhan/gatewaygo/handlers"
	"github.com/minhajuddinkhan/todogo/server"
	nsq "github.com/nsqio/go-nsq"
)

func main() {

	svr := server.NewServer()

	r := mux.NewRouter()

	//host=localhost port=5432 user=dbuser password=dbuser sslmode=disable dbname=dbuser
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=dbuser password=dbuser sslmode=disable dbname=dbuser")
	//	db = db.LogMode(true)
	if err != nil {
		panic(err)
	}

	producer, err := nsq.NewProducer(":4150", nsq.NewConfig())
	if err != nil {
		panic(err)
	}

	r.HandleFunc("/listener", handlers.ListenerHandler(db, producer)).Methods("POST")
	svr.Listen(":3000", r)
}

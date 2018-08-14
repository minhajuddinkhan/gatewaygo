package main

import (
	"github.com/gorilla/mux"
	"github.com/minhajuddinkhan/gatewaygo/handlers"
	"github.com/minhajuddinkhan/gatewaygo/store"
	"github.com/minhajuddinkhan/todogo/server"
	nsq "github.com/nsqio/go-nsq"
)

func main() {

	svr := server.NewServer()

	r := mux.NewRouter()

	pgStore, err := store.NewPostgresStore("host=localhost port=5432 user=dbuser password=dbuser sslmode=disable dbname=dbuser")
	//host=localhost port=5432 user=dbuser password=dbuser sslmode=disable dbname=dbuser
	if err != nil {
		panic(err)
	}

	producer, err := nsq.NewProducer(":4150", nsq.NewConfig())
	if err != nil {
		panic(err)
	}

	r.HandleFunc("/listener", handlers.ListenerHandler(pgStore, producer)).Methods("POST")
	svr.Listen(":3000", r)
}

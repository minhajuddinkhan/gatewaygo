package main

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/minhajuddinkhan/gatewaygo/handlers"
	"github.com/minhajuddinkhan/todogo/server"
)

func main() {

	svr := server.NewServer()

	r := mux.NewRouter()

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=dbuser password=dbuser sslmode=disable dbname=dbuser")
	db = db.LogMode(true)
	if err != nil {
		panic(err)
	}

	r.HandleFunc("/listener", handlers.ListenerHandler(db)).Methods("POST")
	svr.Listen(":3000", r)
}

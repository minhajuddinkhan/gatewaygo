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

	//host=localhost port=5432 user=dbuser password=dbuser sslmode=disable dbname=dbuser
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=admin password=admin sslmode=disable dbname=tododb")
	db = db.LogMode(true)
	if err != nil {
		panic(err)
	}

	r.HandleFunc("/listener", handlers.ListenerHandler(db)).Methods("POST")
	r.HandleFunc("/test", handlers.TestHandler(db)).Methods("POST")
	r.HandleFunc("/mapper", handlers.MapperHandler(db)).Methods("POST")
	svr.Listen(":3000", r)
}

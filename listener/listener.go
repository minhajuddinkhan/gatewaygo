package listener

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/minhajuddinkhan/gatewaygo/handlers"
	"github.com/minhajuddinkhan/todogo/server"
)

func GoListener(port string, connectionStr string) {

	svr := server.NewServer()

	r := mux.NewRouter()

	db, err := gorm.Open("postgres", connectionStr)
	db = db.LogMode(false)
	if err != nil {
		panic(err)
	}

	r.HandleFunc("/listener", handlers.ListenerHandler(db)).Methods("POST")
	svr.Listen(":"+port, r)

}

func main() {

}

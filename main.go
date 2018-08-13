package main

import (
	"github.com/minhajuddinkhan/gatewaygo/listener"
)

func main() {

	listenerPort := "3000"
	connStr := "host=localhost port=5432 user=dbuser password=dbuser sslmode=disable dbname=dbuser"
	listener.GoListener(listenerPort, connStr)
}



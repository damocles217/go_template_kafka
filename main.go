package main

import (
	"fmt"

	"github.com/damocles217/microservice/server"
	"github.com/damocles217/microservice/server/broker"
)

func main() {

	go broker.SubscribeEvents()

	fmt.Println("Starting server...")

	router := server.CreateServer()

	router.Run("localhost:5000")
}

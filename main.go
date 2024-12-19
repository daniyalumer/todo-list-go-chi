package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/daniyalumer/todo-list-go-chi/conf"
)

func main() {
	conf.Setup()

	log.Printf("Starting server on :%v", conf.HttpPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", conf.HttpPort), GetRouter()); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/daniyalumer/todo-list-go-chi/conf"
	"github.com/daniyalumer/todo-list-go-chi/db"
	_ "github.com/lib/pq"
)

//	@title			Todo List API
//	@version		1.0
//	@description	This is a sample server for a todo list application.
//
//	@host			localhost:3000
func main() {
	err := conf.Setup()
	if err != nil {
		log.Fatalf("error loading .env file")
	}

	db.RunMigrations()

	log.Printf("Starting server on :%v", conf.HttpPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", conf.HttpPort), GetRouter()); err != nil {
		log.Fatalf("could not start server: %s\n", err)
	}
}

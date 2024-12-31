package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/daniyalumer/todo-list-go-chi/conf"
	"github.com/daniyalumer/todo-list-go-chi/db"
)

// @title Todo List API
// @version 1.0
// @description This is a sample server for a todo list application.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /api
func main() {
	err := conf.Setup()
	if err != nil {
		log.Fatalf("error loading .env file")
	}

	err = db.Connect()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	db.RunMigrations()

	log.Printf("Starting server on :%v", conf.HttpPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", conf.HttpPort), GetRouter()); err != nil {
		log.Fatalf("could not start server: %s\n", err)
	}
}

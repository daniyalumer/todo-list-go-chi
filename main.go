package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/daniyalumer/todo-list-go-chi/conf"
	"github.com/daniyalumer/todo-list-go-chi/db"
	_ "github.com/lib/pq"
)

func main() {
	err := conf.Setup()
	if err != nil {
		log.Fatalf("error loading .env file")
	}

	err = db.Connect()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Run migrations
	db.RunMigrations()

	// Uncomment the following line to rollback migrations if needed
	// migrations.RollbackMigrations()

	log.Printf("Starting server on :%v", conf.HttpPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", conf.HttpPort), GetRouter()); err != nil {
		log.Fatalf("could not start server: %s\n", err)
	}
}

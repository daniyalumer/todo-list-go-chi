package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/daniyalumer/todo-list-go-chi/internal/handler"
	"github.com/daniyalumer/todo-list-go-chi/internal/routes"

	"github.com/joho/godotenv"
)

func main() {
	r := routes.MainRoute()

	todoRoutes := handler.TodoRoutes()
	r.Mount("/todo", todoRoutes)

	userRoutes := handler.UserRoutes()
	r.Mount("/user", userRoutes)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	HTTP_PORT := os.Getenv("HTTP_PORT")
	log.Printf("Starting server on :%v", HTTP_PORT)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", HTTP_PORT), r); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}

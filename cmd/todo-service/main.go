package main

import (
	"log"
	"net/http"
	"os"

	controller "github.com/daniyalumer/todo-list-go-chi/internal/controller"
	"github.com/joho/godotenv"
)

func main() {
	r := controller.MainRoute()

	todoRoutes := controller.TodoRoutes()
	r.Mount("/todo", todoRoutes)

	userRoutes := controller.UserRoutes()
	r.Mount("/user", userRoutes)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Println("Starting server on :3000")
	if err := http.ListenAndServe(os.Getenv("PORT"), r); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}

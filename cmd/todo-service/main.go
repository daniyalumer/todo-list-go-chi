package main

import (
	"log"
	"net/http"
	"os"

	controller "github.com/daniyalumer/todo-list-go-chi/internal/controller"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Todo app!"))
	})

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

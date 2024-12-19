package main

import (
	"net/http"

	"github.com/daniyalumer/todo-list-go-chi/internal/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func GetRouter() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	todoRoutes := handler.TodoRoutes()
	r.Mount("/todo", todoRoutes)

	userRoutes := handler.UserRoutes()
	r.Mount("/user", userRoutes)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Todo app!"))
	})
	return r
}

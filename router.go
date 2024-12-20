package main

import (
	"github.com/daniyalumer/todo-list-go-chi/internal/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func GetRouter() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", handler.Home)

	r.Route("/api", func(r chi.Router) {
		r.Route("/todo", func(r chi.Router) {
			r.Get("/", handler.GetTodos)
			r.Post("/{userid}", handler.CreateTodo)
			r.Put("/{id}", handler.UpdateTodo)
			r.Delete("/{id}", handler.DeleteTodo)
		})

		r.Route("/user", func(r chi.Router) {
			r.Get("/", handler.GetUser)
			r.Post("/", handler.CreateUser)
			r.Delete("/{userid}", handler.DeleteUser)
		})
	})

	return r
}

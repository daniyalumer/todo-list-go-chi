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
			r.Post("/", handler.CreateTodo)
			r.Put("/", handler.UpdateTodo)
			r.Delete("/", handler.DeleteTodo)
		})

		r.Route("/user", func(r chi.Router) {
			r.Get("/", handler.GetUser)
			r.Post("/", handler.CreateUser)
			r.Delete("/", handler.DeleteUser)
		})
	})

	return r
}

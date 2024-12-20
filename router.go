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
			r.Post("/{user_id}", handler.CreateTodo)
			r.Put("/{todo_id}", handler.UpdateTodo)
			r.Delete("/{todo_id}", handler.DeleteTodo)
		})

		r.Route("/user", func(r chi.Router) {
			r.Get("/", handler.GetUser)
			r.Post("/", handler.CreateUser)
			r.Delete("/{user_id}", handler.DeleteUser)
		})
	})

	return r
}

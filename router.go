package main

import (
	_ "github.com/daniyalumer/todo-list-go-chi/docs"
	"github.com/daniyalumer/todo-list-go-chi/internal/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

func GetRouter() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Route("/api", func(r chi.Router) {
		r.Get("/home", handler.Home)

		r.Route("/todo", func(r chi.Router) {
			r.Get("/", handler.GetTodos)
			r.Post("/user/{user_id}", handler.CreateTodo)
			r.Put("/{todo_id}", handler.UpdateTodo)
			r.Delete("/{todo_id}", handler.DeleteTodo)
		})

		r.Route("/user", func(r chi.Router) {
			r.Get("/", handler.GetUsers)
			r.Post("/", handler.CreateUser)
			r.Delete("/{user_id}", handler.DeleteUser)
		})

		r.Get("/swagger/*", httpSwagger.WrapHandler)
	})

	return r
}

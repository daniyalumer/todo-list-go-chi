package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"log"

	"github.com/daniyalumer/todo-list-go-chi/internal/api"
	"github.com/daniyalumer/todo-list-go-chi/internal/http/rq"
	"github.com/daniyalumer/todo-list-go-chi/internal/service"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := service.ReadTodoList()
	if err != nil {
		log.Printf("error reading todo list: %v", err)
		api.ParseResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	api.ParseResponse(w, todos, http.StatusOK)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var body rq.Todo
	err := api.ParseRequest(r, &body)
	if err != nil {
		log.Printf("error parsing request: %v", err)
		api.ParseResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := api.ParseURLParameter(r, "user_id")
	if err != nil {
		log.Printf("error parsing URL parameter: %v", err)
		api.ParseResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("error converting user_id to integer: %v", err)
		api.ParseResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	todo, err := service.CreateTodo(uint(userID), body)
	if err != nil {
		log.Printf("error creating todo: %v", err)
		api.ParseResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	api.ParseResponse(w, fmt.Sprintf("successfully created todo with id: %d", todo.ID), http.StatusOK)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	var body rq.TodoUpdate
	err := api.ParseRequest(r, &body)
	if err != nil {
		log.Printf("error parsing request: %v", err)
		api.ParseResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := api.ParseURLParameter(r, "todo_id")
	if err != nil {
		log.Printf("error parsing URL parameter: %v", err)
		api.ParseResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	todoID, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("error converting todo_id to integer: %v", err)
		api.ParseResponse(w, fmt.Errorf("unable to process"), http.StatusInternalServerError)
		return
	}

	todo, err := service.UpdateTodo(uint(todoID), body)
	if err != nil {
		log.Printf("error updating todo: %v", err)
		api.ParseResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	api.ParseResponse(w, fmt.Sprintf("successfully updated todo with id: %d", todo.ID), http.StatusOK)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := api.ParseURLParameter(r, "todo_id")
	if err != nil {
		log.Printf("error parsing URL parameter: %v", err)
		api.ParseResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	todoID, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("error converting todo_id to integer: %v", err)
		api.ParseResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	todo, err := service.DeleteTodo(uint(todoID))
	if err != nil {
		log.Printf("error deleting todo: %v", err)
		api.ParseResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	api.ParseResponse(w, fmt.Sprintf("successfully deleted todo with id: %d", todo.ID), http.StatusOK)
}

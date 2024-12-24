package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/daniyalumer/todo-list-go-chi/internal/api"
	"github.com/daniyalumer/todo-list-go-chi/internal/http/rq"
	"github.com/daniyalumer/todo-list-go-chi/internal/service"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := service.ReadTodoList()
	if err != nil {
		api.ParseResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	api.ParseResponse(w, todos, http.StatusOK)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var body rq.TodoCreate
	err := api.ParseRequest(r, &body)
	if err != nil {
		api.ParseResponse(w, err.Error(), http.StatusInternalServerError)
	}

	id, err := api.ParseURLParameter(r, "user_id")
	if err != nil {
		api.ParseResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(id)
	if err != nil {
		api.ParseResponse(w, fmt.Errorf("unable to process"), http.StatusInternalServerError)
		return
	}

	todo, err := service.CreateTodo(uint(userID), body)
	if err != nil {
		api.ParseResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	api.ParseResponse(w, fmt.Sprintf("successfully created todo with id: %d", todo.ID), http.StatusOK)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	var body rq.TodoUpdate
	err := api.ParseRequest(r, &body)
	if err != nil {
		api.ParseResponse(w, err.Error(), http.StatusInternalServerError)
	}

	id, err := api.ParseURLParameter(r, "todo_id")
	if err != nil {
		api.ParseResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	todoID, err := strconv.Atoi(id)
	if err != nil {
		api.ParseResponse(w, fmt.Errorf("unable to process"), http.StatusInternalServerError)
		return
	}

	todo, err := service.UpdateTodo(uint(todoID), body)
	if err != nil {
		api.ParseResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	api.ParseResponse(w, fmt.Sprintf("successfully updated todo with id: %d", todo.ID), http.StatusOK)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := api.ParseURLParameter(r, "todo_id")
	if err != nil {
		api.ParseResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	todoID, err := strconv.Atoi(id)
	if err != nil {
		api.ParseResponse(w, fmt.Errorf("unable to process"), http.StatusInternalServerError)
		return
	}

	todo, err := service.DeleteTodo(uint(todoID))
	if err != nil {
		api.ParseResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	api.ParseResponse(w, fmt.Sprintf("successfully deleted todo with id: %d", todo.ID), http.StatusOK)
}

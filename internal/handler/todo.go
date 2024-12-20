package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/daniyalumer/todo-list-go-chi/internal/api"
	"github.com/daniyalumer/todo-list-go-chi/internal/models"
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
	var body models.TodoCreateRequest
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

	todo, err := service.CreateTodo(userID, body)
	if err != nil {
		api.ParseResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	api.ParseResponse(w, todo, http.StatusOK)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	var body models.TodoUpdateRequest
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

	todo, err := service.UpdateTodo(todoID, body)
	if err != nil {
		api.ParseResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	api.ParseResponse(w, todo, http.StatusOK)
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

	todo, err := service.DeleteTodo(todoID)
	if err != nil {
		api.ParseResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	api.ParseResponse(w, todo, http.StatusOK)
}

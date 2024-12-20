package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/daniyalumer/todo-list-go-chi/internal/api"
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
	todoDescription := r.Form.Get("description")

	id, err := api.ParseURLParameter(r, "userId")
	if err != nil {
		api.ParseResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(id)
	if err != nil {
		api.ParseResponse(w, fmt.Errorf("unable to process"), http.StatusInternalServerError)
		return
	}

	todo, err := service.CreateTodo(todoDescription, userID)
	if err != nil {
		api.ParseResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	api.ParseResponse(w, todo, http.StatusOK)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	todoDescription := r.Form.Get("description")
	completedStr := r.Form.Get("completed")

	Completed := completedStr == "true"

	id, err := api.ParseURLParameter(r, "id")
	if err != nil {
		api.ParseResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	todoID, err := strconv.Atoi(id)
	if err != nil {
		api.ParseResponse(w, fmt.Errorf("unable to process"), http.StatusInternalServerError)
		return
	}

	todo, err := service.UpdateTodo(todoID, Completed, todoDescription)
	if err != nil {
		api.ParseResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	api.ParseResponse(w, todo, http.StatusOK)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := api.ParseURLParameter(r, "id")
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

package handler

import (
	"net/http"
	"strconv"

	"github.com/daniyalumer/todo-list-go-chi/internal/api"
	"github.com/daniyalumer/todo-list-go-chi/internal/service"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := service.ReadTodoList()
	if err != nil {
		api.ResponseWriter(w, err.Error(), http.StatusBadRequest)
		return
	}
	api.ResponseWriter(w, todos, http.StatusOK)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	if !api.ParseForm(w, r) {
		return
	}

	todoDescription := r.Form.Get("description")

	userID, err := api.ParseURLParameter(r, "userId")
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusInternalServerError)
		return
	}

	todo, err := service.CreateTodo(todoDescription, id)
	if err != nil {
		api.ResponseWriter(w, err.Error(), http.StatusInternalServerError)
		return
	}

	api.ResponseWriter(w, todo, http.StatusOK)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	if !api.ParseForm(w, r) {
		return
	}

	todoDescription := r.Form.Get("description")
	completedStr := r.Form.Get("completed")

	Completed := completedStr == "true"

	idstr := r.URL.Query().Get("id")
	Id, err := api.ParseURLParameter(idstr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	todo, err := service.UpdateTodo(Id, Completed, todoDescription)
	if err != nil {
		api.ResponseWriter(w, err.Error(), http.StatusBadRequest)
		return
	}
	api.ResponseWriter(w, todo, http.StatusOK)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	if !api.ParseForm(w, r) {
		return
	}

	idstr := r.URL.Query().Get("id")

	Id, err := api.ParseURLParameter(idstr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	todo, err := service.DeleteTodo(Id)
	if err != nil {
		api.ResponseWriter(w, err.Error(), http.StatusBadRequest)
		return
	}
	api.ResponseWriter(w, todo, http.StatusOK)
}

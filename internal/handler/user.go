package handler

import (
	"fmt"
	"net/http"

	"github.com/daniyalumer/todo-list-go-chi/internal/api"
	"github.com/daniyalumer/todo-list-go-chi/internal/service"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	users, err := service.ReadUsers()
	if err != nil {
		api.ResponseWriter(w, err.Error(), http.StatusBadRequest)
		return
	}
	api.ResponseWriter(w, users, http.StatusOK)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	_, id, err := service.CreateUser()
	if err != nil {
		api.ResponseWriter(w, err.Error(), http.StatusBadRequest)
		return
	}
	api.ResponseWriter(w, fmt.Sprintf("successfully deleted user with id: %d", id), http.StatusOK)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if !api.ParseForm(w, r) {
		return
	}

	idstr := r.URL.Query().Get("id")

	id, err := api.ParseURLParameter(idstr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	_, err = service.DeleteUser(id)
	if err != nil {
		api.ResponseWriter(w, err.Error(), http.StatusBadRequest)
		return
	}

	api.ResponseWriter(w, fmt.Sprintf("successfully deleted user with id: %d", id), http.StatusOK)
}

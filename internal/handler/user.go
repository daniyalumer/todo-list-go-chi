package handler

import (
	"fmt"
	"net/http"

	"github.com/daniyalumer/todo-list-go-chi/internal/helper"
	"github.com/daniyalumer/todo-list-go-chi/internal/service"
	"github.com/go-chi/chi/v5"
)

func UserRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", GetUserHandler)
	r.Post("/", CreateUserHandler)
	r.Delete("/", DeleteUserHandler)

	return r
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	users, err := service.ReadUsers()
	if err != nil {
		helper.ResponseWriter(w, err.Error(), http.StatusBadRequest)
		return
	}
	helper.ResponseWriter(w, users, http.StatusOK)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	_, id, err := service.CreateUser()
	if err != nil {
		helper.ResponseWriter(w, err.Error(), http.StatusBadRequest)
		return
	}
	helper.ResponseWriter(w, fmt.Sprintf("successfully deleted user with id: %d", id), http.StatusOK)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	if !helper.ParseForm(w, r) {
		return
	}

	idstr := r.URL.Query().Get("id")

	id, err := helper.ConvertIdToInteger(idstr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	_, err = service.DeleteUser(id)
	if err != nil {
		helper.ResponseWriter(w, err.Error(), http.StatusBadRequest)
		return
	}

	helper.ResponseWriter(w, fmt.Sprintf("successfully deleted user with id: %d", id), http.StatusOK)
}

package controller

import (
	"net/http"
	"strconv"

	"github.com/daniyalumer/todo-list-go-chi/internal/helper"
	"github.com/daniyalumer/todo-list-go-chi/internal/models"
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
		helper.UserListResponseWriter(w, "failed to read user list", []models.User{}, http.StatusBadRequest)
		return
	}
	helper.UserListResponseWriter(w, "successfully read user list", users, http.StatusAccepted)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	user, err := service.CreateUser()
	if err != nil {
		helper.UserResponseWriter(w, "failed to read user list", models.User{}, http.StatusBadRequest)
		return
	}
	helper.UserResponseWriter(w, "successfully created user", user, http.StatusAccepted)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	if !helper.ParseForm(w, r) {
		return
	}

	IdStr := r.URL.Query().Get("id")

	Id, err := strconv.Atoi(IdStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	user, err := service.DeleteUser(Id)
	if err != nil {
		helper.UserResponseWriter(w, "failed to read user list", models.User{}, http.StatusBadRequest)
		return
	}
	helper.UserResponseWriter(w, "successfully deleted user", user, http.StatusAccepted)
}

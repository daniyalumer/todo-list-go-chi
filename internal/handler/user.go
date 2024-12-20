package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/daniyalumer/todo-list-go-chi/internal/api"
	"github.com/daniyalumer/todo-list-go-chi/internal/service"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	users, err := service.ReadUsers()
	if err != nil {
		api.ParseResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	api.ParseResponse(w, users, http.StatusOK)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user, err := service.CreateUser()
	if err != nil {
		api.ParseResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	api.ParseResponse(w, fmt.Sprintf("successfully created user with id: %d", user.ID), http.StatusOK)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := api.ParseURLParameter(r, "userid")
	if err != nil {
		api.ParseResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(id)
	if err != nil {
		api.ParseResponse(w, fmt.Errorf("unable to process"), http.StatusInternalServerError)
		return
	}

	_, err = service.DeleteUser(userID)
	if err != nil {
		api.ParseResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	api.ParseResponse(w, fmt.Sprintf("successfully deleted user with id: %d", userID), http.StatusOK)
}

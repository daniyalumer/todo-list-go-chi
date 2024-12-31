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

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := service.ReadUsers()
	if err != nil {
		log.Printf("Error reading users: %v", err)
		api.ParseResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	api.ParseResponse(w, users, http.StatusOK)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var body rq.User
	err := api.ParseRequest(r, &body)
	if err != nil {
		log.Printf("Error parsing request: %v", err)
		api.ParseResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := service.CreateUser(body)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		api.ParseResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	api.ParseResponse(w, fmt.Sprintf("successfully created user with id: %d", user.ID), http.StatusOK)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := api.ParseURLParameter(r, "user_id")
	if err != nil {
		log.Printf("Error parsing URL parameter: %v", err)
		api.ParseResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("Error converting user ID to integer: %v", err)
		api.ParseResponse(w, fmt.Errorf("unable to process"), http.StatusInternalServerError)
		return
	}

	_, err = service.DeleteUser(uint(userID))
	if err != nil {
		log.Printf("Error deleting user: %v", err)
		api.ParseResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	api.ParseResponse(w, fmt.Sprintf("successfully deleted user with id: %d", userID), http.StatusOK)
}

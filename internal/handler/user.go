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

// GetUsers godoc
//
//	@Summary		Get all users
//	@Description	Get all users
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		rq.User
//	@Failure		400	{string}	string	"Bad Request"
//	@Router			/user [get]
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := service.ReadUsers()
	if err != nil {
		log.Printf("Error reading users: %v", err)
		api.ParseResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	api.ParseResponse(w, users, http.StatusOK)
}

// CreateUser godoc
//
//	@Summary		Create a new user
//	@Description	Create a new user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		rq.User	true	"User"
//	@Success		200		{string}	string	"Successfully created user"
//	@Failure		400		{string}	string	"Bad Request"
//	@Failure		500		{string}	string	"Internal Server Error"
//	@Router			/user [post]
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

// DeleteUser godoc
//
//	@Summary		Delete a user
//	@Description	Delete a user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user_id	path		int		true	"User ID"
//	@Success		200		{string}	string	"Successfully deleted user"
//	@Failure		400		{string}	string	"Bad Request"
//	@Failure		500		{string}	string	"Internal Server Error"
//	@Router			/user/{user_id} [delete]
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

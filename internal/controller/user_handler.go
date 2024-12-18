package todo

import (
	"encoding/json"
	"net/http"
	"strconv"

	models "github.com/daniyalumer/todo-list-go-chi/internal/models"
	service "github.com/daniyalumer/todo-list-go-chi/internal/service"
	"github.com/go-chi/chi/v5"
)

func UserRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/get", GetUserHandler)
	r.Post("/post", CreateUserHandler)
	r.Delete("/delete", DeleteUserHandler)

	return r
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	message, users := service.ReadUsers()
	w.Header().Set("Content-Type", "application/json")

	response := struct {
		Message string        `json:"message"`
		Users   []models.User `json:"todos"`
	}{
		Message: message,
		Users:   users,
	}

	json.NewEncoder(w).Encode(response)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	message, user := service.CreateUser()
	w.Header().Set("Content-Type", "application/json")

	response := struct {
		Message string      `json:"message"`
		User    models.User `json:"todo"`
	}{
		Message: message,
		User:    user,
	}
	json.NewEncoder(w).Encode(response)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
	}

	IdStr := r.FormValue("id")

	Id, err := strconv.Atoi(IdStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	message, user := service.DeleteUser(Id)
	w.Header().Set("Content-Type", "application/json")

	response := struct {
		Message string      `json:"message"`
		User    models.User `json:"todo"`
	}{
		Message: message,
		User:    user,
	}
	json.NewEncoder(w).Encode(response)
}

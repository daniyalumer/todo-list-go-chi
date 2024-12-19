package helper

import (
	"encoding/json"
	"net/http"

	"github.com/daniyalumer/todo-list-go-chi/internal/models"
)

func TodoResponseWriter(w http.ResponseWriter, message string, todo models.Todo, statusCode uint) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(statusCode))

	response := struct {
		Message string      `json:"message"`
		ToDo    models.Todo `json:"todo"`
	}{
		Message: message,
		ToDo:    todo,
	}
	json.NewEncoder(w).Encode(response)
}

func TodoListResponseWriter(w http.ResponseWriter, message string, todos []models.Todo, statusCode uint) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(statusCode))

	response := struct {
		Message string        `json:"message"`
		ToDo    []models.Todo `json:"todo"`
		
	}{
		Message: message,
		ToDo:    todos,
	}
	json.NewEncoder(w).Encode(response)
}

func UserResponseWriter(w http.ResponseWriter, message string, user models.User, statusCode uint) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(statusCode))

	response := struct {
		Message string      `json:"message"`
		User    models.User `json:"todo"`
	}{
		Message: message,
		User:    user,
	}
	json.NewEncoder(w).Encode(response)
}

func UserListResponseWriter(w http.ResponseWriter, message string, users []models.User, statusCode uint) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(statusCode))

	response := struct {
		Message string        `json:"message"`
		Users   []models.User `json:"todos"`
	}{
		Message: message,
		Users:   users,
	}

	json.NewEncoder(w).Encode(response)
}

func ParseForm(w http.ResponseWriter, r *http.Request) bool {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return false
	}
	return true
}

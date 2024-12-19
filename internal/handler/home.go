package handler

import (
	"net/http"

	"github.com/daniyalumer/todo-list-go-chi/internal/api"
)

func Home(w http.ResponseWriter, r *http.Request) {
	api.ResponseWriter(w, "Welcome to Todo app!", http.StatusOK)
}

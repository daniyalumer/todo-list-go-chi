package handler

import (
	"net/http"

	"github.com/daniyalumer/todo-list-go-chi/internal/api"
)

// Home godoc
// @Summary Home endpoint
// @Description Welcome message for the Todo app
// @Tags home
// @Accept  json
// @Produce  json
// @Success 200 {string} string "Welcome to Todo app!"
// @Router /home [get]
func Home(w http.ResponseWriter, r *http.Request) {
	api.ParseResponse(w, "Welcome to Todo app!", http.StatusOK)
}

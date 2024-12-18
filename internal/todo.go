package todo

import (
	"time"

	"github.com/go-chi/chi/v5"
)

func NewRouter() chi.Router {
	return TodoRoutes() 
}

type Todo struct {
	ID            int       `json:"id"`
	Description   string    `json:"description"`
	DateCreated   time.Time `json:"dateCreated"`
	DateUpdated   time.Time `json:"dateUpdated"`
	DateCompleted time.Time `json:"dateCompleted"`
	Completed     bool      `json:"completed"`
}

var todoList []Todo
var nextID = 1

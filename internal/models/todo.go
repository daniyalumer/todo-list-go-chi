package todo

import (
	"time"
)

type Todo struct {
	ID            int       `json:"id"`
	Description   string    `json:"description"`
	DateCreated   time.Time `json:"date_created"`
	DateUpdated   time.Time `json:"date_updated"`
	DateCompleted time.Time `json:"date_completed"`
	Completed     bool      `json:"completed"`
	UserID        int       `json:"user_id"`
}

var TodoList []Todo
var TodoID = 1

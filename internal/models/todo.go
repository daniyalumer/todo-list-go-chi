package models

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

type TodoCreateRequest struct {
	Description string `json:"description"`
}

type TodoUpdateRequest struct {
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

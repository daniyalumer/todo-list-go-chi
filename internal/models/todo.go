package todo

import (
	"time"
)

type Todo struct {
	ID            int       `json:"id"`
	Description   string    `json:"description"`
	DateCreated   time.Time `json:"dateCreated"`
	DateUpdated   time.Time `json:"dateUpdated"`
	DateCompleted time.Time `json:"dateCompleted"`
	Completed     bool      `json:"completed"`
	UserID        int       `json:"userId"`
}

var TodoList []Todo
var TodoID = 1

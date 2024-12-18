package todo

type User struct {
	ID    int    `json:"ID"`
	Todos []Todo `json:"todos"`
}

var UserList []User
var UserId = 1

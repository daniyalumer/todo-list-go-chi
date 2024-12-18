package todo

import (
	"log"

	models "github.com/daniyalumer/todo-list-go-chi/internal/models"
)

func CreateUser() (string, models.User) {
	newUser := models.User{
		ID:    models.TodoID,
		Todos: []models.Todo{},
	}
	models.TodoID++
	models.UserList = append(models.UserList, newUser)
	return "User Added Successfully", newUser
}

func ReadUsers() (string, []models.User) {
	return "User Read Successfully", models.UserList
}

func DeleteUser(userid int) (string, models.User) {
	for index, User := range models.UserList {
		if User.ID == userid {
			models.UserList = append(models.UserList[:index], models.UserList[index+1:]...)
			for index, Todo := range models.TodoList {
				if Todo.UserID == userid {
					models.TodoList = append(models.TodoList[:index], models.TodoList[index+1:]...)
					log.Printf("User Items Id: %d And Description: %v Removed Successfully", Todo.ID, Todo.Description)
				}
			}
			return "User Deleted Successfully", User
		}
	}
	return "User Not Found", models.User{}
}

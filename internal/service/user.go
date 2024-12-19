package service

import (
	"fmt"
	"log"

	"github.com/daniyalumer/todo-list-go-chi/internal/models"
)

func CreateUser() (models.User, error) {
	newUser := models.User{
		ID:    models.UserId,
		Todos: []models.Todo{},
	}
	models.UserId++
	models.UserList = append(models.UserList, newUser)
	return newUser, nil
}

func ReadUsers() ([]models.User, error) {
	if len(models.TodoList) == 0 {
		return nil, fmt.Errorf("user list empty")
	}
	return models.UserList, nil
}

func DeleteUser(userid int) (models.User, error) {
	for index, User := range models.UserList {
		if User.ID == userid {
			models.UserList = append(models.UserList[:index], models.UserList[index+1:]...)
			for index, Todo := range models.TodoList {
				if Todo.UserID == userid {
					models.TodoList = append(models.TodoList[:index], models.TodoList[index+1:]...)
					log.Printf("User Items Id: %d And Description: %v Removed Successfully", Todo.ID, Todo.Description)
				}
			}
			return User, nil
		}
	}
	return models.User{}, fmt.Errorf("user not found to delete")
}

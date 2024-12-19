package service

import (
	"fmt"
	"log"
	"time"

	"github.com/daniyalumer/todo-list-go-chi/internal/models"
)

func CreateTodo(description string, userid int) (models.Todo, error) {

	userExists := false
	for _, User := range models.UserList {
		if User.ID == userid {
			userExists = true
			break
		}
	}

	if !userExists {
		return models.Todo{}, fmt.Errorf("provided user does not exist")
	}

	newTodo := models.Todo{
		ID:          models.TodoID,
		Description: description,
		DateCreated: time.Now(),
		DateUpdated: time.Time{},
		Completed:   false,
		UserID:      userid,
	}
	models.TodoID++
	models.TodoList = append(models.TodoList, newTodo)
	models.UserList[userid-1].Todos = append(models.UserList[userid-1].Todos, newTodo)
	return newTodo, nil
}

func ReadTodoList() ([]models.Todo, error) {
	if len(models.TodoList) == 0 {
		return nil, fmt.Errorf("user list empty")
	}
	return models.TodoList, nil
}

func UpdateTodo(id int, completed bool, description string) (models.Todo, error) {
	if (description == "" && !completed) || (description != "" && completed) {
		return models.Todo{}, fmt.Errorf("provide either description or completed")
	}

	if description != "" {
		for index, todoItem := range models.TodoList {
			if todoItem.ID == id {
				currentTime := time.Now()
				models.TodoList[index].Description = description
				models.TodoList[index].DateUpdated = currentTime
				for index1, User := range models.UserList {
					for index2, Todo := range User.Todos {
						if Todo.ID == id {
							models.UserList[index1].Todos[index2].Description = description
							models.UserList[index1].Todos[index2].DateUpdated = currentTime
							log.Printf("Description: %v Updated Successfully In User Todos For Todo Id: %d", Todo.Description, Todo.ID)
						}
					}
				}
				return models.TodoList[index], nil
			}
		}
		return models.Todo{}, fmt.Errorf("item not found to update description")
	}

	if completed {
		for index, todoItem := range models.TodoList {
			fmt.Println(todoItem.ID)
			if todoItem.ID == id {
				if models.TodoList[index].Completed {
					return models.TodoList[index], fmt.Errorf("todo item already marked completed")
				}
				currentTime := time.Now()
				models.TodoList[index].Completed = completed
				models.TodoList[index].DateUpdated = currentTime
				for index1, User := range models.UserList {
					for index2, Todo := range User.Todos {
						if Todo.ID == id {
							models.UserList[index1].Todos[index2].Completed = completed
							models.UserList[index1].Todos[index2].DateUpdated = currentTime
							log.Printf("Description: %v Updated Successfully In User Todos For Todo Id: %d", Todo.Description, Todo.ID)
						}
					}
				}
				return models.TodoList[index], nil
			}
		}
	}
	return models.Todo{}, fmt.Errorf("todo item not found to mark completed")
}

func DeleteTodo(id int) (models.Todo, error) {
	for index, todoItem := range models.TodoList {
		if todoItem.ID == id {
			models.TodoList = append(models.TodoList[:index], models.TodoList[index+1:]...)
			for index1, User := range models.UserList {
				for index2, Todo := range User.Todos {
					if Todo.ID == id {
						models.UserList[index1].Todos = append(models.UserList[index1].Todos[:index2], models.UserList[index1].Todos[index2+1:]...)
						log.Printf("User Todos Id: %d And Description: %v Removed Successfully From User Todos", Todo.ID, Todo.Description)
					}
				}
			}
			return todoItem, nil
		}
	}
	return models.Todo{}, fmt.Errorf("todo item not found to delete")
}

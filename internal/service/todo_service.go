package todo

import (
	"fmt"
	"log"
	"time"

	models "github.com/daniyalumer/todo-list-go-chi/internal/models"
)

func CreateTodo(description string, userid int) (string, models.Todo) {
	for _, User := range models.UserList {
		if User.ID != userid {
			return "User Does Not Exist", models.Todo{}
		}
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
	models.UserList[userid].Todos = append(models.UserList[userid].Todos, newTodo)
	return "Todo Created Successfully", newTodo
}

func ReadTodoList() (string, []models.Todo) {
	return "TodoList Read Successfully", models.TodoList
}

func UpdateTodo(id int, completed bool, description string) (string, models.Todo) {
	if description != "" {
		for index, todoItem := range models.TodoList {
			fmt.Println(todoItem.ID)
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
				return "Todo Item Updated Successfully In TodoList", models.TodoList[index]
			}
		}
		return "Item Not Found To Update Description", models.Todo{}
	}

	if completed {
		for index, todoItem := range models.TodoList {
			fmt.Println(todoItem.ID)
			if todoItem.ID == id {
				if models.TodoList[index].Completed {
					return "Todo Item Already Marked Compleded", models.TodoList[index]
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
				return "Todo Item Marked Completed Successfully", models.TodoList[index]
			}
		}
	}
	return "Item Not Found To Mark Completed", models.Todo{}
}

func DeleteTodo(id int) (string, models.Todo) {
	for index, todoItem := range models.TodoList {
		if todoItem.ID == id {
			models.TodoList = append(models.TodoList[:index], models.TodoList[index+1:]...)
			for index1, User := range models.UserList {
				for index2, Todo := range User.Todos {
					if Todo.ID == id {
						models.UserList[index1].Todos = append(models.UserList[index].Todos[:index2], models.UserList[index].Todos[index2+1:]...)
						log.Printf("User Todos Id: %d And Description: %v Removed Successfully From User Todos", Todo.ID, Todo.Description)
					}
				}
			}
			return "Item Deleted Successfully From TodoList", todoItem
		}
	}
	return "Item Not Found", models.Todo{}
}

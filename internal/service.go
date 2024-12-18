package todo

import (
	"fmt"
	"time"
)

func ReadTodoList() (string, []Todo) {
	return "TodoList Read Successfully", todoList
}

func CreateTodo(description string) (string, Todo) {
	newTodo := Todo{
		ID:          nextID,
		Description: description,
		DateCreated: time.Now(),
		DateUpdated: time.Time{},
		Completed:   false,
	}
	nextID++
	todoList = append(todoList, newTodo)
	return "Todo Created Successfully", newTodo
}

func UpdateTodo(id int, completed bool, description string) (string, Todo) {
	if description != "" {
		for index, todoItem := range todoList {
			fmt.Println(todoItem.ID)
			if todoItem.ID == id {
				todoList[index].Description = description
				todoList[index].DateUpdated = time.Now()
				return "Todo Item Updated Successfully", todoList[index]
			}
		}
		return "Item Not Found To Update Description", Todo{}
	}

	if completed {
		for index, todoItem := range todoList {
			fmt.Println(todoItem.ID)
			if todoItem.ID == id {
				if todoList[index].Completed {
					return "Todo Item Already Marked Compleded", todoList[index]
				}
				todoList[index].Completed = completed
				todoList[index].DateUpdated = time.Now()
				return "Todo Item Marked Completed Successfully", todoList[index]
			}
		}
	}
	return "Item Not Found To Mark Completed", Todo{}
}

func DeleteTodo(id int) (string, Todo) {
	for index, todoItem := range todoList {
		if todoItem.ID == id {
			todoList = append(todoList[:index], todoList[index+1:]...)
			return "Item Deleted Successfully", todoItem
		}
	}
	return "Item Not Found", Todo{}
}

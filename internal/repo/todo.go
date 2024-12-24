package repository

import (
	"github.com/daniyalumer/todo-list-go-chi/db"
	"github.com/daniyalumer/todo-list-go-chi/internal/models"
)

func Create(todo *models.Todo) error {
	result := db.DB.Create(&todo)
	return result.Error
}

func FindAll(todos *[]models.Todo) error {
	result := db.DB.Find(&todos)
	return result.Error
}

func FindById(todo *models.Todo, todoID uint) error {
	result := db.DB.First(&todo, todoID)
	return result.Error
}

func Update(todo *models.Todo, updates interface{}) error {
	results := db.DB.Model(&todo).Updates(updates)
	return results.Error
}

func Delete(todo *models.Todo) error {

	if err := db.DB.Delete(&todo).Error; err != nil {
		return err
	}

	return nil

}

package repo

import (
	"github.com/daniyalumer/todo-list-go-chi/db/dao"
	"gorm.io/gorm"
)

func Create(todo *dao.Todo, DB *gorm.DB) error {
	return DB.Create(&todo).Error
}

func FindAll(todos *[]dao.Todo, DB *gorm.DB) error {
	return DB.Find(todos).Error
}

func FindById(todo *dao.Todo, DB *gorm.DB, todoID uint) error {
	return DB.First(&todo, todoID).Error
}

func Update(todo *dao.Todo, DB *gorm.DB, updates interface{}) error {
	return DB.Model(&todo).Updates(updates).Error
}

func Delete(todo *dao.Todo, DB *gorm.DB) error {
	return DB.Delete(&todo).Error
}

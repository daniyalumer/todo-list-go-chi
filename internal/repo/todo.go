package repo

import (
	"github.com/daniyalumer/todo-list-go-chi/db"
	"github.com/daniyalumer/todo-list-go-chi/db/dao"
)

func Create(todo *dao.Todo) error {
	DB := db.GetConnection()
	return DB.Create(&todo).Error
}

func FindAll(todos *[]dao.Todo) error {
	DB := db.GetConnection()
	return DB.Find(todos).Error
}

func FindById(todo *dao.Todo, todoID uint) error {
	DB := db.GetConnection()
	return DB.First(&todo, todoID).Error
}

func Update(todo *dao.Todo, updates interface{}) error {
	DB := db.GetConnection()
	return DB.Model(&todo).Updates(updates).Error
}

func Delete(todo *dao.Todo) error {
	DB := db.GetConnection()
	return DB.Delete(&todo).Error
}

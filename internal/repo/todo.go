package repo

import (
	"context"

	"github.com/daniyalumer/todo-list-go-chi/db"
	"github.com/daniyalumer/todo-list-go-chi/db/dao"
)

func Create(ctx context.Context, todo *dao.Todo) error {
	DB := db.GetConnection()
	return DB.WithContext(ctx).Create(&todo).Error
}

func FindAll(ctx context.Context, todos *[]dao.Todo) error {
	DB := db.GetConnection()
	return DB.WithContext(ctx).Find(todos).Error
}

func FindById(ctx context.Context, todo *dao.Todo, todoID uint) error {
	DB := db.GetConnection()
	return DB.WithContext(ctx).First(&todo, todoID).Error
}

func Update(ctx context.Context, todo *dao.Todo, updates interface{}) error {
	DB := db.GetConnection()
	return DB.WithContext(ctx).Model(&todo).Updates(updates).Error
}

func Delete(ctx context.Context, todo *dao.Todo) error {
	DB := db.GetConnection()
	return DB.WithContext(ctx).Delete(&todo).Error
}

package repo

import (
	"context"

	"github.com/daniyalumer/todo-list-go-chi/db"
	"github.com/daniyalumer/todo-list-go-chi/db/dao"
)

func Create(ctx context.Context, todo *dao.Todo) error {
	return db.Conn().WithContext(ctx).Create(&todo).Error
}

func FindAll(ctx context.Context, todos *[]dao.Todo) error {
	return db.Conn().WithContext(ctx).Find(todos).Error
}

func FindById(ctx context.Context, todo *dao.Todo, todoID uint) error {
	return db.Conn().WithContext(ctx).First(&todo, todoID).Error
}

func Update(ctx context.Context, todo *dao.Todo, updates interface{}) error {
	return db.Conn().WithContext(ctx).Model(&todo).Updates(updates).Error
}

func Delete(ctx context.Context, todo *dao.Todo) error {
	return db.Conn().WithContext(ctx).Delete(&todo).Error
}

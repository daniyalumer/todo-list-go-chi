package repo

import (
	"context"

	"github.com/daniyalumer/todo-list-go-chi/db"
	"github.com/daniyalumer/todo-list-go-chi/db/dao"
)

func CreateUser(ctx context.Context, user *dao.User) error {
	DB := db.Conn()
	return DB.WithContext(ctx).Create(&user).Error
}

func FindAllUsers(ctx context.Context, users *[]dao.User) error {
	DB := db.Conn()
	return DB.WithContext(ctx).Preload("Todos").Find(&users).Error
}

func FindByIdUser(ctx context.Context, user *dao.User, userID uint) error {
	DB := db.Conn()
	return DB.WithContext(ctx).Preload("Todos").Find(&user, userID).Error
}

func DeleteUser(ctx context.Context, user *dao.User) error {
	DB := db.Conn()
	if err := DB.WithContext(ctx).Where("user_id = ?", user.ID).Delete(&dao.Todo{}).Error; err != nil {
		return err
	}

	if err := DB.WithContext(ctx).Delete(&user).Error; err != nil {
		return err
	}

	return nil
}

func CheckDeleted(ctx context.Context, user *dao.User, userID uint) error {
	DB := db.Conn()
	return DB.WithContext(ctx).Where("deleted_at IS NULL").Preload("Todos").First(user, userID).Error
}

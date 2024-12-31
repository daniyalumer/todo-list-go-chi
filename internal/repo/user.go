package repository

import (
	"github.com/daniyalumer/todo-list-go-chi/db"
	"github.com/daniyalumer/todo-list-go-chi/db/dao"
	"gorm.io/gorm"
)

func CreateUser(user *dao.User) error {
	result := db.DB.Create(&user)
	return result.Error
}

func FindAllUsers(users *[]dao.User) error {
	result := db.DB.Preload("Todos").Find(&users)
	return result.Error
}

func FindByIdUser(user *dao.User, userID uint) error {
	result := db.DB.Preload("Todos").Find(&user, userID)
	return result.Error
}

func DeleteUser(tx *gorm.DB, user *dao.User) error {
	if err := tx.Where("user_id = ?", user.ID).Delete(&dao.Todo{}).Error; err != nil {
		return err
	}

	if err := tx.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}

func CheckDeleted(user *dao.User, userID uint) error {
	result := db.DB.Where("deleted_at IS NULL").Preload("Todos").First(user, userID)
	return result.Error
}

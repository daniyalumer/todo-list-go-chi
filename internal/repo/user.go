package repo

import (
	"github.com/daniyalumer/todo-list-go-chi/db/dao"
	"gorm.io/gorm"
)

func CreateUser(user *dao.User, DB *gorm.DB) error {
	return DB.Create(&user).Error
}

func FindAllUsers(users *[]dao.User, DB *gorm.DB) error {
	return DB.Preload("Todos").Find(&users).Error
}

func FindByIdUser(user *dao.User, DB *gorm.DB, userID uint) error {
	return DB.Preload("Todos").Find(&user, userID).Error
}

func DeleteUser(DB *gorm.DB, user *dao.User) error {
	if err := DB.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}

func CheckDeleted(user *dao.User, DB *gorm.DB, userID uint) error {
	return DB.Where("deleted_at IS NULL").Preload("Todos").First(user, userID).Error
}

package db

import (
	"fmt"
	"log"

	"github.com/daniyalumer/todo-list-go-chi/conf"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	var DB *gorm.DB

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		conf.DbHost,
		conf.DbPort,
		conf.DbUser,
		conf.DbName,
		conf.DbPassword,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	DB = db
	log.Println("Database connected successfully")

	return DB, err
}

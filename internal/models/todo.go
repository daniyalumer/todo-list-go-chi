package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Description string     `json:"description" gorm:"not null"`
	CompletedAt *time.Time `json:"completed_at" gorm:"default:null"`
	Completed   bool       `json:"completed" gorm:"default:false"`
	UserID      uint       `json:"user_id" gorm:"not null"`
}

package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Mail struct {
	gorm.Model
	ID      uuid.UUID `gorm:"type:uuid;default:UUID()" `
	Name    string    `json:"name"`
	Image   string    `json:"image"`
	Content string    `json:"content"`
	Date    string    `json:"date"`
	Sender  string    `json:"sender"`
}

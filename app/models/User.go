package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       uuid.UUID `json:"id" gorm:"type:uuid;default:UUID()" validate:"required,uuid"`
	Name     string    `json:"name" validate:"required,lte=255"`
	Email    string    `json:"email" validate:"required,lte=255" gorm:"unique"`
	Password []byte    `json:"-"`
}

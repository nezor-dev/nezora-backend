package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Bookmark struct {
	gorm.Model
	ID   uuid.UUID `gorm:"type:uuid;default:UUID()" validate:"required,uuid"`
	Name string    `json:"name" validate:"required,lte=255"`
	Url  string    `json:"url" validate:"required,lte=255"`
}

package models

import "gorm.io/gorm"

type Document struct {
	gorm.Model
	Name    string `json:"name"`
	Image   string `json:"image"`
	Content string `json:"content"`
	Date    string `json:"date"`
	Sender  string `json:"sender"`
}

package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title string `gorm:"unique" json:"title,omitempty" bson:"title" binding:"required,min=5"`
	Body  string `json:"body,omitempty" bson:"body" binding:"required"`
}

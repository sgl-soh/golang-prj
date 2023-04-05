package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique" json:"email,omitempty" bson:"email" binding:"required"`
	Password string `json:"password,omitempty" bson:"password" binding:"required"`
}

package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Token    string `gorm:"-"`
	Blog     []Blog `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type UserResponse struct {
	Message string
	Data    []User
}

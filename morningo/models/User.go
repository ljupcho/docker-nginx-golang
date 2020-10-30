package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	FirstName string `json:"first_name" form:"first_name"`
	LastName string `json:"last_name" form:"last_name"`
	Email string `json:"email" form:"email"`
	Age int `json:"age" form:"age"`
}

func (User) TableName() string {
	return "users"
}

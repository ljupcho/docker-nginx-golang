package models

import "github.com/jinzhu/gorm"
// import "time"

type User struct {
	gorm.Model
	FirstName string `json:"first_name" form:"first_name"`
	LastName string `json:"last_name" form:"last_name"`
	Email string `json:"email" form:"email"`
	Age int `json:"age" form:"age"`
	// CreatedAt time.Time `json:"created_at" form:"created_at"`
	// UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

func (User) TableName() string {
	return "users"
}

package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	FirstName string `json:"first_name" form:"first_name"`
	LastName string `json:"last_name" form:"last_name"`
	Email string `gorm:"type:varchar(255);unique_index" form:"email"`
	Age int `json:"age" form:"age"`
	GroupID uint `json:"group_id" form:"group_id"`
	Group Group  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Posts []Post `gorm:"foreignKey:UserID"`
}

func (User) TableName() string {
	return "users"
}

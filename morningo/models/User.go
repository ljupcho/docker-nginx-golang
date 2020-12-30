package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	FirstName string `gorm:"not null" json:"first_name" form:"first_name"`
	LastName string `gorm:"not null" json:"last_name" form:"last_name"`
	Email string `gorm:"type:varchar(255);unique_index;not null" form:"email"`
	Age int `gorm:"not null" json:"age" form:"age"`
	GroupID uint `gorm:"not null" json:"group_id" form:"group_id"`
	Group Group  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Posts []Post `gorm:"foreignKey:UserID"`
}

func (User) TableName() string {
	return "users"
}

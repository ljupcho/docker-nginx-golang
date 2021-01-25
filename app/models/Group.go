package models

import "github.com/jinzhu/gorm"

type Group struct {
	gorm.Model
	Name string `gorm:"not null" json:"name" form:"name"`
	Phone string `gorm:"not null" json:"phone" form:"phone"`
	Email string `gorm:"type:varchar(255);unique_index;not null" json:"email" form:"email"`
	City string `gorm:"not null" json:"city" form:"city"`	
}

func (Group) TableName() string {
	return "groups"
}
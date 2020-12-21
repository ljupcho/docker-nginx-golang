package models

import "github.com/jinzhu/gorm"

type Group struct {
	gorm.Model
	Name string `json:"name" form:"name"`
	Phone string `json:"phone" form:"phone"`
	Email string `gorm:"type:varchar(255);unique_index" json:"email" form:"email"`
	City string `json:"city" form:"city"`	
}

func (Group) TableName() string {
	return "groups"
}
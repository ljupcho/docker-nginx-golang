package models

import "github.com/jinzhu/gorm"

type Post struct {
	gorm.Model
	Title string `gorm:"not null" json:"title" form:"title"`
	Content string `gorm:"not null" json:"content" form:"content"`
	UserID uint `gorm:"not null" form:"user_id" json:"user_id"`
}

func (Post) TableName() string {
	return "posts"
}

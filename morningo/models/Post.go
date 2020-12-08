package models

import "github.com/jinzhu/gorm"
// import "time"

type Post struct {
	gorm.Model
	Title string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	UserId uint `form:"user_id" json:"user_id"`
}

func (Post) TableName() string {
	return "posts"
}

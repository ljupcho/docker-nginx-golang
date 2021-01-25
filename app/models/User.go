package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FirstName string `gorm:"not null" json:"first_name" form:"first_name"`
	LastName string `gorm:"not null" json:"last_name" form:"last_name"`
	Email string `gorm:"type:varchar(255);unique_index;not null" json:"email" form:"email"`
	Password string `gorm:"type:varchar(255);not null" json: "password" form:"password"`
	Gender string `gorm:"type:varchar(50)" json:"gender" form:"gender"`
}

func (User) TableName() string {
	return "users"
}
	
func AddUser(user *User) {
	Model.Create(&user)
	return
}

func UserById(userId uint) (user User) {
	Model.Where("id = ?", userId).First(&user)
	return
}

func UserByName(name string) (user User) {
    Model.Where("name = ?", name).First(&user)
    return
}

func UserByEmail(email string) (user User) {
    Model.Where("email = ?", email).First(&user)
    return
}
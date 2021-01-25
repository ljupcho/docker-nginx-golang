package models

import (
	"github.com/jinzhu/gorm"
	"app/config"
)

var Model *gorm.DB

func init() {
	var err error
	Model, err = gorm.Open("mysql", config.GetEnv().Database.FormatDSN())

	Model.LogMode(true)
	Model.DropTableIfExists(User{})
	Model.LogMode(true).AutoMigrate(&User{})

	if err != nil {
		panic(err)
	}
}

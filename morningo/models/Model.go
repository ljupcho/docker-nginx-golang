package models

import (
	"github.com/jinzhu/gorm"
	"morningo/config"
	"morningo/modules/log"
)

var Model *gorm.DB

func init() {
	var err error
	log.Println(config.GetEnv().Database.FormatDSN())
	Model, err = gorm.Open("mysql", config.GetEnv().Database.FormatDSN())

	// Model.DropTableIfExists(Post{}, User{})
	// Model.LogMode(true).AutoMigrate(&User{}, &Post{})
	// Model.LogMode(true).Model(&Post{}).AddForeignKey("user_id", "users(user_id)", "CASCADE", "CASCADE")

	if err != nil {
		panic(err)
	}
}

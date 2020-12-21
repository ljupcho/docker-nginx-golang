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

	Model.LogMode(true)
	// Model.DropTableIfExists(Group{}, Post{}, User{})
	// Model.LogMode(true).AutoMigrate(&Group{}, &User{}, &Post{})
	// Model.LogMode(true).Model(&User{}).AddForeignKey("group_id", "groups(id)", "CASCADE", "CASCADE")
	// Model.LogMode(true).Model(&Post{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

	if err != nil {
		panic(err)
	}
}

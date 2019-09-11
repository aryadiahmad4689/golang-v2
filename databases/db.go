package databases

import (
	"fmt"
	"v2/config"
	"v2/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

// Init ..
func Init() {
	configuration := config.GetConfig()
	ConnectString := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", configuration.DbUsername, configuration.DbPassword, configuration.DbName)
	db, err = gorm.Open("mysql", ConnectString)
	if err != nil {
		panic("DB Connection Error")
	}
	db.AutoMigrate(&models.Posts{})

}

// DbManager ..
func DbManager() *gorm.DB {
	// defer db.Close()

	return db
}

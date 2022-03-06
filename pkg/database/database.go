package database

import (
	"log"

	"go-just-portfolio/models"
	config "go-just-portfolio/pkg/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitDB() *gorm.DB {
	conf := config.GetConfig()
	db, err := gorm.Open("mysql", conf.MYSQL)
	if err != nil {
		log.Panicln(err)
	}

	migrateDB(db)
	return db
}

func migrateDB(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Project{})
	db.AutoMigrate(&models.Photo{})
	db.AutoMigrate(&models.Tags{})
	db.AutoMigrate(&models.Description{})
	db.AutoMigrate(&models.Category{})
}

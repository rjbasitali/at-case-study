package db

import (
	"git.rjbasitali.com/at-case-study/cfg"
	"git.rjbasitali.com/at-case-study/pkg/log"
	"git.rjbasitali.com/at-case-study/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	db, err := gorm.Open(postgres.Open(cfg.DBConnStr), &gorm.Config{})

	if err != nil {
		log.Error(err)
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Product{})
	DB = db
}

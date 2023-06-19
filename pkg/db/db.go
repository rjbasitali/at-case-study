package db

import (
	"fmt"

	"git.rjbasitali.com/at-case-study/pkg/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(connStr string) {
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		fmt.Fprintln(gin.DefaultErrorWriter, "error connecting to database: ", err)
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Product{})
	DB = db
}

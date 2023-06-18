package db

import (
	"fmt"

	"git.rjbasitali.com/at-case-study/cfg"
	"git.rjbasitali.com/at-case-study/pkg/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	db, err := gorm.Open(postgres.Open(cfg.DB_CONN_STR), &gorm.Config{})

	if err != nil {
		fmt.Fprintln(gin.DefaultErrorWriter, "error connecting to database: ", err)
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Product{})
	DB = db
}

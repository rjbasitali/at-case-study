package db

import (
	"fmt"

	"git.rjbasitali.com/at-case-study/pkg/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Init initializes the database connection.
// It accepts the connection string as a parameter.
// It panics if the connection to the database could not be established.
func Init(connStr string) {
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		fmt.Fprintln(gin.DefaultErrorWriter, "error connecting to database: ", err)
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Product{})
	DB = db
}

package db

import (
	"fmt"
	"time"

	"git.rjbasitali.com/at-case-study/pkg/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Init initializes the database connection.
// It accepts the connection string as a parameter.
// It panics if the connection to the database could not be established.
// It also sets the maximum idle connections, maximum open connections and connection max lifetime.
func Init(connStr string) {
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		fmt.Fprintln(gin.DefaultErrorWriter, "error connecting to database: ", err)
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Product{})

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Fprintln(gin.DefaultErrorWriter, "error getting database connection: ", err)
		panic("failed to get database connection")
	}

	err = sqlDB.Ping()
	if err != nil {
		fmt.Fprintln(gin.DefaultErrorWriter, "error pinging database: ", err)
		panic("failed to ping database")
	}
	fmt.Fprintln(gin.DefaultWriter, "Connected to database")

	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(200)
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = db
}

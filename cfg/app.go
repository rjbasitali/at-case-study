package cfg

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

var (
	APP_ADDR string

	JWT_SIGNING_SERCRET string
	JWT_EXPIRY_DURATION time.Duration

	DB_NAME string
	DB_USER string
	DB_PASS string
	DB_HOST string

	DB_CONN_STR string

	REDIS_HOST     string
	REDIS_PORT     string
	REDIS_PASSWORD string
	REDIS_DB       int
)

// Init initializes the application configuration.
// It reads the configuration from the env file.
// It should be called before any other function.
func Init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	APP_ADDR = os.Getenv("APP_ADDR")

	JWT_SIGNING_SERCRET = os.Getenv("JWT_SIGNING_SERCRET")
	JWT_EXPIRY_DURATION, _ = time.ParseDuration(os.Getenv("JWT_EXPIRY_DURATION"))

	DB_NAME = os.Getenv("DB_NAME")
	DB_USER = os.Getenv("DB_USER")
	DB_PASS = os.Getenv("DB_PASS")
	DB_HOST = os.Getenv("DB_HOST")

	REDIS_HOST = os.Getenv("REDIS_HOST")
	REDIS_PORT = os.Getenv("REDIS_PORT")
	REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")
	REDIS_DB, _ = strconv.Atoi(os.Getenv("REDIS_DB"))

	DB_CONN_STR = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", DB_USER, DB_PASS, DB_HOST, DB_NAME)
}

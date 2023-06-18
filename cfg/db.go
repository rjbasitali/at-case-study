package cfg

import (
	"fmt"
	"os"
)

var (
	DB_NAME string
	DB_USER string
	DB_PASS string
	DB_HOST string
)

var (
	DBConnStr = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", DB_USER, DB_PASS, DB_HOST, DB_NAME)
)

func init() {
	if val, ok := os.LookupEnv("DB_NAME"); ok {
		DB_NAME = val
	}

	if val, ok := os.LookupEnv("DB_USER"); ok {
		DB_USER = val
	}

	if val, ok := os.LookupEnv("DB_PASS"); ok {
		DB_PASS = val
	}

	if val, ok := os.LookupEnv("DB_HOST"); ok {
		DB_HOST = val
	}
}

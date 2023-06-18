package cfg

import (
	"os"
	"strconv"
	"time"
)

var (
	APP_ADDR string

	JWT_SIGNING_SERCRET string
	JWT_EXPIRY_DURATION time.Duration

	REDIS_HOST     string
	REDIS_PORT     string
	REDIS_PASSWORD string
	REDIS_DB       int
)

func init() {
	if val, ok := os.LookupEnv("APP_ADDR"); ok {
		APP_ADDR = val
	}

	if val, ok := os.LookupEnv("JWT_SIGNING_SERCRET"); ok {
		JWT_SIGNING_SERCRET = val
	}

	if val, ok := os.LookupEnv("JWT_EXPIRY_DURATION"); ok {
		if d, err := time.ParseDuration(val); err == nil {
			JWT_EXPIRY_DURATION = d
		}
	}

	if val, ok := os.LookupEnv("REDIS_HOST"); ok {
		REDIS_HOST = val
	}

	if val, ok := os.LookupEnv("REDIS_PORT"); ok {
		REDIS_PORT = val
	}

	if val, ok := os.LookupEnv("REDIS_PASSWORD"); ok {
		REDIS_PASSWORD = val
	}

	if val, ok := os.LookupEnv("REDIS_DB"); ok {
		if i, err := strconv.Atoi(val); err == nil {
			REDIS_DB = i
		}
	}
}

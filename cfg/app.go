package cfg

import "time"

const (
	APP_ADDR = ":8080"

	JWT_SIGNING_SERCRET = "OY2/L9am5+A6MlIAGCF53HqH2NZ8o5IaR3AyveskrJyh+7Z2J2PdeXAxbE4dfQhA"
	JWT_EXPIRY_DURATION = time.Duration(30) * time.Minute

	REDIS_ADDR     = "localhost:6379"
	REDIS_PASSWORD = ""
	REDIS_DB       = 0
)

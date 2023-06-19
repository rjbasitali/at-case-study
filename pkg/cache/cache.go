package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var c *redis.Client

func Init(host, port, password string, db int) {
	c = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       db,
	})

	_, err := c.Ping(c.Context()).Result()
	if err != nil {
		fmt.Fprintln(gin.DefaultErrorWriter, "error while connecting to redis", err)
		panic(err)
	}
}

func Set(ctx context.Context, key string, value interface{}) error {
	p, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.Set(ctx, key, p, 0).Err()
}

func Get(ctx context.Context, key string, dest interface{}) error {
	val, err := c.Get(ctx, key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(val), dest)
}

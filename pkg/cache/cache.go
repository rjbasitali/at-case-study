package cache

import (
	"context"
	"encoding/json"

	"git.rjbasitali.com/at-case-study/cfg"
	"git.rjbasitali.com/at-case-study/pkg/log"
	"github.com/go-redis/redis/v8"
)

var c *redis.Client

func init() {
	c = redis.NewClient(&redis.Options{
		Addr:     cfg.REDIS_ADDR,
		Password: cfg.REDIS_PASSWORD,
		DB:       cfg.REDIS_DB,
	})

	_, err := c.Ping(c.Context()).Result()
	if err != nil {
		log.Error("error while connecting to redis", err)
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

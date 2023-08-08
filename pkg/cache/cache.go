package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type cache struct {
	Client *redis.Client
	sync.Mutex
}

var c = &cache{}

// Init initializes the redis client.
// It accepts the host, port, password and db as parameters.
// It panics if the connection to redis could not be established.
func Init(host, port, password string, db int) {
	if c.Client != nil {
		return
	}

	c.Lock()
	defer c.Unlock()

	if c.Client == nil {
		c.Client = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", host, port),
			Password: password,
			DB:       db,
		})

		_, err := c.Client.Ping(c.Client.Context()).Result()
		if err != nil {
			fmt.Fprintln(gin.DefaultErrorWriter, "error while connecting to redis", err)
			panic(err)
		}
	}
}

// Set sets a key value pair in redis.
// It accepts a context, key and value as parameters.
// It returns an error if the key value pair could not be set.
func Set(ctx context.Context, key string, value interface{}) error {
	p, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.Client.Set(ctx, key, p, 0).Err()
}

// Get gets a value from redis.
// It accepts a context and key as parameters.
// It returns an error if the value could not be fetched.
func Get(ctx context.Context, key string, dest interface{}) error {
	val, err := c.Client.Get(ctx, key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(val), dest)
}

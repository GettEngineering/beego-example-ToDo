package store

import (
	"time"

	"gopkg.in/redis.v4"
)

var StoreClient *redis.Client

func init() {
	StoreClient = redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		Password:     "", // no password set
		MaxRetries:   3,
		PoolTimeout:  time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
		DialTimeout:  2 * time.Second,
	})
}

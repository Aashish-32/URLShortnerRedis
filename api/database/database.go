package database

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

func CreateClient(dbno int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       1,
	})
	return rdb
}

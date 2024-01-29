package redis

import (
	"context"
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	rdb *redis.Client

	Init = sync.OnceFunc(setRDB)
)

func setRDB() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := rdb.Ping(context.TODO()).Result()
	if err != nil {
		panic(err)
	}
}

func DB() *redis.Client {
	return rdb
}

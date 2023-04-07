package redisCli

import (
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

func NewRedis(host string, port string) *redis.Redis {
	redisdb, err := redis.NewRedis(redis.RedisConf{
		Type: "none",
		Host: "localhost:6379",
	})
	if err != nil {
		fmt.Printf("Failed to connect to redis")
		panic(err)
	}

	return redisdb
}

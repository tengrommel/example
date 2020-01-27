package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

// redis

var redisDb *redis.Client

func initRedis() error {
	redisDb = redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379", Password: "", DB: 0})
	_, err := redisDb.Ping().Result()
	return err
}

func main() {
	err := initRedis()
	if err != nil {
		fmt.Printf("connect redis failed, err: %v\n", err)
		return
	}
	fmt.Println("连接成功")

	key := "rank"
	items := []redis.Z{
		redis.Z{Score: 90, Member: "PHP"},
		redis.Z{Score: 96, Member: "Go"},
	}
	redisDb.ZAdd(key, items...)
}

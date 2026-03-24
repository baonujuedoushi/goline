package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func redisSample() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 地址
		Password: "",               // 密码，没有则留空
		DB:       0,                // 使用的数据库
		PoolSize: 10,               // 连接池大小
	})

	ctx := context.Background()
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("连接成功:", pong)
}

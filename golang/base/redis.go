package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func redisSample() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 地址
		Password: "",               // 密码，没有则留空
		DB:       0,                // 使用的数据库
		PoolSize: 10,               
	})
	defer rdb.Close()

	ctx := context.Background()
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("连接成功:", pong)

	userID := "1001"
	userKey := fmt.Sprintf("user:%s", userID)

	rdb.HSet(ctx, userKey, map[string]any{
		"name": "test user",
		"age":  25,
	})
	rdb.Expire(ctx, userKey, 1*time.Hour)

	user, _ := rdb.HGetAll(ctx, userKey).Result()
	fmt.Printf("user info: %+v\n", user)

	visitKey := fmt.Sprintf("visits:%s", userID)
	count, _ := rdb.Incr(ctx, visitKey).Result()
	fmt.Printf("visitis counts: %d\n", count)

	pipe := rdb.Pipeline()
	for i := 0; i < 10; i++ {
		pipe.SAdd(ctx, "visited_users", fmt.Sprintf("user_%d", i))
	}
	pipe.Exec(ctx)

	users, _ := rdb.SMembers(ctx, "visited_users").Result()
	fmt.Printf("has visited: %v\n", users)
}

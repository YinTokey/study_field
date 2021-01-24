package pkg

import (
	"fmt"
	"github.com/spf13/viper"

	"strconv"

	"github.com/go-redis/redis"
)

// RedisClient Redis缓存客户端单例
var RedisClient *redis.Client

// Redis 在中间件中初始化redis链接
func Redis() {

	db, _ := strconv.ParseUint(viper.GetString("redis.db"), 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:       viper.GetString("redis.addr"),
		Password:   viper.GetString("redis.pw"),
		DB:         int(db),
		MaxRetries: 1,
	})

	_, err := client.Ping().Result()

	if err != nil {
		Log().Panic("连接Redis不成功", err)
	}

	fmt.Println("Redis 连接成功")

	RedisClient = client
}

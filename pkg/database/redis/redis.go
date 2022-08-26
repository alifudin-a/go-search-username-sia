package database

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// Untuk SIA
var RCSIA *redis.Client

// Untuk SIA
func OpenRedisSIA() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Username: "",
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       2,
	})

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		log.Println("[SIA] An error occured while sending PING to redis : ", err)
	}

	log.Println("[SIA] Redis connection =====> PING :", pong)

	RCSIA = client
}

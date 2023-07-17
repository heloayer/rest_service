package initialize

import (
	"fmt"
	"log"
	"os"


	"github.com/go-redis/redis"
)

var Rdb *redis.Client

func Init() error {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	if redisHost == "" || redisPort == "" { // нужен docker-compose up для подключения к Redis
		return fmt.Errorf("Отсутствуют обязательные переменные окружения для подключения к Redis") 
	}

	Rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: redisPassword,
		DB:       0,
	})

	if _, err := Rdb.Ping().Result(); err != nil {
		return fmt.Errorf("failed to connect to Redis: %v", err)
	}

	log.Println("Connected to Redis")
	return nil
}

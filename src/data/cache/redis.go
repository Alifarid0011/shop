package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/Alifarid0011/shop/src/config"
	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func InitRedis(cfg *config.Config) error {
	redisClient = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password:     cfg.Redis.Password,
		DB:           0,
		DialTimeout:  cfg.Redis.DialTimeout * time.Second,
		WriteTimeout: cfg.Redis.WriteTimeout * time.Second,
		ReadTimeout:  cfg.Redis.ReadTimeout * time.Second,
		PoolTimeout:  cfg.Redis.PoolTimeout * time.Second,
	})
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	return nil
}

func GetRedis() *redis.Client {
	return redisClient
}
func CloseRedis() {
	err := redisClient.Close()
	if err != nil {
		return
	}
}

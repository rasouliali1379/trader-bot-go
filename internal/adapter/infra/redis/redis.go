package redis

import (
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"hamgit.ir/novin-backend/trader-bot/config"
	"time"
)

func Init(cfg *config.Config) *redis.Client {
	zap.L().Info("Establishing connection with redis db")
	defer zap.L().Info("Connection with redis db established successfully")
	return redis.NewClient(&redis.Options{
		Addr:        config.C().Redis.Host,
		Password:    config.C().Redis.Password,
		DB:          config.C().Redis.DB,
		ReadTimeout: time.Duration(config.C().Redis.Timeout) * time.Second,
	})
}

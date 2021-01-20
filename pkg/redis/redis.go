package redis

import (
	"github.com/go-redis/redis/v8"
	"github.com/jxlwqq/go-restful/internal/config"
)

type Redis struct {
	*redis.Client
}

func New(cfg *config.Config) *Redis  {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddress,
	})

	return &Redis{client}
}
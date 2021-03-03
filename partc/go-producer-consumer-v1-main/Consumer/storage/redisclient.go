package storage

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var (
	client = &RedisClient{}
)

type RedisClient struct {
	c      *redis.Client
	logger *log.Logger
}

package cache

import (
	"time"
	"github.com/gomodule/redigo/redis"
)

// Wraps the Redis client to meet the Cache interface.
type RedisStore struct {
	pool              *redis.Pool
	defaultExpiration time.Duration
}
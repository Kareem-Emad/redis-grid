package cache

import "github.com/gomodule/redigo/redis"

// Client a contract interface for any caching third party client to be used with our system
type Client interface {
	Init(string) error
	InitWithDefault() error
	ExecuteRawCommand(commandName string, args ...interface{}) (string, error)
}

// RedisClient implementation for the cache client interface using redis
type RedisClient struct {
	redisConnection redis.Conn
}

package cache

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

const tag = "[cache-client:redis]"

// InitWithDefault intialize a connection to redis server
func (rc *RedisClient) InitWithDefault() error {
	conn, err := redis.Dial("tcp", redisConnectionURL)
	rc.redisConnection = conn

	return err
}

// Init intialize a connection to redis server
func (rc *RedisClient) Init(connURL string) error {
	conn, err := redis.Dial("tcp", connURL)
	rc.redisConnection = conn

	return err
}

// ExecuteRawCommand sends a command string to redis
func (rc *RedisClient) ExecuteRawCommand(commandName string, args ...interface{}) (string, error) {
	log.Printf("%s executing command %s %s", tag, commandName, args)
	result, err := rc.redisConnection.Do(commandName, args...)

	res, err := redis.String(result, err)

	log.Printf("%s execution result %s | error", tag, res)
	return res, err
}

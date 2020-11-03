package cache

import (
	"github.com/Kareem-Emad/redis-grid/envreader"
)

var redisConnectionURL = envreader.ReadEnvAsString("REDIS_CONNECTION_URL", ":6379")

package distributer

import (
	"github.com/Kareem-Emad/redis-grid/cache"
	"github.com/serialx/hashring"
)

// ShardsManager responsible for distributing and reading data from appropiate shards
type ShardsManager struct {
	connectedShards map[string]cache.Client
	shardsRing      *hashring.HashRing
}

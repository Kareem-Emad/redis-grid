package distributer

import (
	"fmt"
	"log"
	"strings"

	"github.com/Kareem-Emad/redis-grid/cache"
	"github.com/serialx/hashring"
)

const tag = "[distributer]"

// InitShardsWithDefault establishes a connection with each cache instance underneath this proxy to
// distribute load on all shards
func (sm *ShardsManager) InitShardsWithDefault() {
	connectionURLs := strings.Split(shardsConnectionsString, ",")
	sm.shardsRing = hashring.New(connectionURLs)

	sm.connectedShards = make(map[string]cache.Client, len(connectionURLs))

	for _, connURL := range connectionURLs {
		currCache := cache.RedisClient{}
		err := currCache.Init(connURL)

		if err != nil {
			log.Fatalf("%s Failed to connect to one of the shards in initialization | error %s", tag, err)
		}

		sm.connectedShards[connURL] = &currCache
	}
	log.Printf("%s shards initialized succesffuly ", tag)
}

// InitShards initalizes the shards mapping and the distribution ring
func (sm *ShardsManager) InitShards(shardsMap map[string]cache.Client, shardIDs []string) {
	sm.connectedShards = shardsMap
	sm.shardsRing = hashring.New(shardIDs)
}

// ExecuteCommand retrieves the appropiate shard to execute the comand on assigned key
func (sm *ShardsManager) ExecuteCommand(key string, commandName string, args ...interface{}) (string, error) {
	shard, exists := sm.shardsRing.GetNode(key)

	if !exists {
		return "", fmt.Errorf("not enough shards to store key")
	}
	log.Printf("%s executing command on key %s for shard %s", tag, key, shard)

	return sm.connectedShards[shard].ExecuteRawCommand(commandName, args...)
}

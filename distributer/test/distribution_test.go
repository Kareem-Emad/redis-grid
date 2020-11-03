package distributertest

import (
	"testing"

	"github.com/Kareem-Emad/redis-grid/distributer"
)

// TestRandomShardAssignment should choose the corresponding shard according to ring mapping
func TestRandomShardAssignment(t *testing.T) {
	var shardman distributer.ShardsManager

	shardman.InitShards(fakeShardsMap, fakeConnectionURLs)

	testKey := "some_key"
	shardman.ExecuteCommand(testKey, "SET", "X", "Y")

	if globalExecutionCount != 1 {
		t.Fatalf("Expected  1 transaction to be executed | found %d", globalExecutionCount)
	}
}

// TestRandomShardAssignmentIntegration should be able to set and fetch the value assigned
func TestRandomShardAssignmentIntegration(t *testing.T) {
	t.Skipf("skipping integration tests | re-enable if redis or a cache server is up")
	var shardman distributer.ShardsManager

	shardman.InitShardsWithDefault()

	testKey := "some_key"
	_, err := shardman.ExecuteCommand(testKey, "SET", testKey, "Y")

	if err != nil {
		t.Fatalf("expected set command to work normally, error encountered: %s", err)
	}

	res, err := shardman.ExecuteCommand(testKey, "GET", testKey)
	if err != nil {
		t.Fatalf("expected get command to work normally, error encountered: %s", err)
	}
	if res != "Y" {
		t.Fatalf("expected value of key %s to be Y | found %s", testKey, res)
	}
}

//TODO: add more tests

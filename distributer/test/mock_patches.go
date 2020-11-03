package distributertest

import "github.com/Kareem-Emad/redis-grid/cache"

var fakeConnectionURLs = []string{
	"http://redis_test1:4000",
	"http://redis_test2:4000",
	"http://redis_test3:4000",
	"http://redis_test4:4000",
}

var globalExecutionCount = 0

var shardExecutionStack = map[string]string{
	fakeConnectionURLs[0]: "",
	fakeConnectionURLs[1]: "",
	fakeConnectionURLs[2]: "",
	fakeConnectionURLs[3]: "",
}

var fakeShardsMap = map[string]cache.Client{
	fakeConnectionURLs[0]: &mockCache{ExecutionCount: 0, GlobalExecutionCount: &globalExecutionCount},
	fakeConnectionURLs[1]: &mockCache{ExecutionCount: 0, GlobalExecutionCount: &globalExecutionCount},
	fakeConnectionURLs[2]: &mockCache{ExecutionCount: 0, GlobalExecutionCount: &globalExecutionCount},
	fakeConnectionURLs[3]: &mockCache{ExecutionCount: 0, GlobalExecutionCount: &globalExecutionCount},
}

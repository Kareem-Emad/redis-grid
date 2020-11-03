package distributer

import (
	"github.com/Kareem-Emad/redis-grid/envreader"
)

var shardsConnectionsString = envreader.ReadEnvAsString("SHARDS_CONNECTION_URLS", ":6379")

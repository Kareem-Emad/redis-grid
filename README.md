# Redis Grid

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Build Status:](https://github.com/Kareem-Emad/redis-grid/workflows/Build/badge.svg)](https://github.com/Kareem-Emad/redis-grid/actions)
[![GoReportCard example](https://goreportcard.com/badge/github.com/Kareem-Emad/redis-grid)](https://goreportcard.com/report/Kareem-Emad/redis-grid)

A basic implemenation for distributed caching layer aimed primarily for redis server to support sharding data and distributing load on multiple servers

## setup
To run server

```shell
make run
```

## environment variables

List of envs needs to be setup before starting the service

- `REDIS_CONNECTION_URL` default url for redis to connect to if none supplied
- `SHARDS_CONNECTION_URLS` comma separated list of redis urls to use for sharding the data
- `SERVER_PORT` port used to serve http requests

## how to use

sample request to send when server is up to do any kind of command on redis

```shell
curl --location --request POST 'http://localhost:5000/execute_command' \
--header 'Content-Type: application/json' \
--data-raw '{
    "command_name":"SET",
    "key": "XY",
    "command_args": [
        "XY",
        "XZ"
    ]
}'
```
Notice you need to mention the key twice, one time in the `command args` array and another time in the `key` field


## can I use another cache server?

yes as long you implemnent the interface `cache.client`, then you can easily swap in your new caching layer instead of redis, no assumptions are made for redis beyond the interface level.
```go
// Client a contract interface for any caching third party client to be used with our system
type Client interface {
	Init(string) error
	InitWithDefault() error
	ExecuteRawCommand(commandName string, args ...interface{}) (string, error)
}
```


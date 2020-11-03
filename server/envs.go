package server

import "github.com/Kareem-Emad/redis-grid/envreader"

var port = envreader.ReadEnvAsString("SERVER_PORT", "5000")

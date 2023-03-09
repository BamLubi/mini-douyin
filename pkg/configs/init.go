package configs

import (
	"math/rand"
	"time"
)

func InitConfigs(logpath string, loglevel string) {
	rand.Seed(time.Now().UnixNano())

	InitLoggerOnly(logpath, loglevel)
	InitDB()
	InitRedis()
	initNacos()
	initMQ()
}

func InitLoggerOnly(path string, level string) {
	initLogger(path, level)
}

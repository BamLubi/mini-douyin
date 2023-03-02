package configs

import (
	"github.com/gomodule/redigo/redis"
)

var RD redis.Conn

func initRedis() {
	var err error
	RD, err = redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		Logger.Error("Redis Connect Fail" + err.Error())
		panic(err)
	}
	Logger.Info("Redis Connect Success.")
}

package main

import (
	"mini-douyin/pkg/configs"
)

func main() {
	configs.InitLoggerOnly("/root/mini-douyin/logs/test.log", "debug")
	configs.InitRedis()

	configs.RD.Send("SELECT", 10)
	configs.RD.Send("set", "test_name", "123123")
	configs.RD.Flush()
	// v, _ := configs.RD.Receive()
	// v, err := configs.RD.Receive()
	// if err != nil {
	// 	log.Println(err)
	// }
	// if v == nil {
	// 	log.Println("空")
	// }
}

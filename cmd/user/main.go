package main

import (
	userdouyin "mini-douyin/cmd/user/kitex_gen/userdouyin/userservice"
	config "mini-douyin/pkg/configs"
)

func Init() {
	config.InitConfigs("/root/mini-douyin/logs/user-rpc.log", "debug")
}

func main() {
	Init()

	svr := userdouyin.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		config.Logger.Error(err.Error())
	}
}

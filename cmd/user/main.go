package main

import (
	userdouyin "mini-douyin/cmd/user/kitex_gen/userdouyin/userservice"
	config "mini-douyin/pkg/configs"
	"net"

	"github.com/cloudwego/kitex/server"
)

func Init() {
	config.InitConfigs("/root/mini-douyin/logs/user-rpc.log", "debug")
}

func main() {
	Init()

	addr, _ := net.ResolveTCPAddr("tcp", "0.0.0.0:8888")
	var opts []server.Option
	opts = append(opts, server.WithServiceAddr(addr))

	svr := userdouyin.NewServer(new(UserServiceImpl), opts...)

	err := svr.Run()

	if err != nil {
		config.Logger.Error(err.Error())
	}
}

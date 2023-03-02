package main

import (
	videodouyin "mini-douyin/cmd/video/kitex_gen/videodouyin/videoservice"
	config "mini-douyin/pkg/configs"
	"net"

	"github.com/cloudwego/kitex/server"
)

func Init() {
	config.InitConfigs("/root/mini-douyin/logs/video-rpc.log", "debug")
}

func main() {
	Init()

	addr, _ := net.ResolveTCPAddr("tcp", "0.0.0.0:9999")
	var opts []server.Option
	opts = append(opts, server.WithServiceAddr(addr))

	svr := videodouyin.NewServer(new(VideoServiceImpl), opts...)

	err := svr.Run()

	if err != nil {
		config.Logger.Error(err.Error())
	}
}

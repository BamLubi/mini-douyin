package main

import (
	socitydouyin "mini-douyin/cmd/socity/kitex_gen/socitydouyin/socityservice"
	"mini-douyin/cmd/socity/rpc"
	config "mini-douyin/pkg/configs"
	"mini-douyin/pkg/consts"
	"mini-douyin/pkg/utils"
	"net"
	"strconv"

	"github.com/cloudwego/kitex/server"
)

var Ip string
var Port uint64

func Init() (err error) {
	config.InitConfigs("/root/mini-douyin/logs/socity-rpc.log", "debug")

	rpc.InitRPC()

	// 获取本机外网ip和端口
	Ip, err = utils.GetIp()
	if err != nil {
		config.Logger.Error(err.Error())
		return
	}
	Port, err = utils.GetFreePort()
	if err != nil {
		config.Logger.Error(err.Error())
		return
	}
	// 服务中心注册服务
	err = config.NacosRegister(Ip, Port, consts.SocityServiceName)
	if err != nil {
		config.Logger.Error(err.Error())
		return
	}
	return err
}

func main() {
	var err error
	err = Init()
	if err != nil {
		panic(err)
	}

	addr, _ := net.ResolveTCPAddr("tcp", Ip+":"+strconv.Itoa(int(Port)))
	var opts []server.Option
	opts = append(opts, server.WithServiceAddr(addr))
	svr := socitydouyin.NewServer(new(SocityServiceImpl), opts...)
	err = svr.Run()
	if err != nil {
		config.Logger.Error(err.Error())
		panic(err)
	}
}

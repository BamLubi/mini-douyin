package rpc

import (
	"mini-douyin/cmd/socity/kitex_gen/userdouyin/userservice"
	config "mini-douyin/pkg/configs"
	"mini-douyin/pkg/consts"
	"strconv"

	"github.com/cloudwego/kitex/client"
)

var UserClient userservice.Client

func initUser() {
	ip, port, err := config.NacosSelectOneHealthyInstance(consts.UserServiceName)
	if err != nil {
		panic(err)
	}
	c, err := userservice.NewClient(consts.UserServiceName, client.WithHostPorts(ip+":"+strconv.Itoa(int(port))))
	if err != nil {
		panic(err)
	}
	UserClient = c
	config.Logger.Info("Init User RPC Client Success!")
}

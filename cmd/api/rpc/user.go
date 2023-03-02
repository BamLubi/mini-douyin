package rpc

import (
	"mini-douyin/cmd/api/kitex_gen/userdouyin/userservice"
	config "mini-douyin/pkg/configs"
	"mini-douyin/pkg/consts"

	"github.com/cloudwego/kitex/client"
)

var UserClient userservice.Client

func initUser() {
	c, err := userservice.NewClient(consts.UserServiceName, client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		panic(err)
	}
	UserClient = c
	config.Logger.Info("Init User RPC Client Success!")
}

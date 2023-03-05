package rpc

import (
	"mini-douyin/cmd/api/kitex_gen/socitydouyin/socityservice"
	config "mini-douyin/pkg/configs"
	"mini-douyin/pkg/consts"
	"strconv"

	"github.com/cloudwego/kitex/client"
)

var SocityClient socityservice.Client

func initSocity() {
	ip, port, err := config.NacosSelectOneHealthyInstance(consts.SocityServiceName)
	if err != nil {
		panic(err)
	}
	c, err := socityservice.NewClient(consts.UserServiceName, client.WithHostPorts(ip+":"+strconv.Itoa(int(port))))
	if err != nil {
		panic(err)
	}
	SocityClient = c
	config.Logger.Info("Init Socity RPC Client Success!")
}

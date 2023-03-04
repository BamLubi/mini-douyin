package rpc

import (
	"mini-douyin/cmd/api/kitex_gen/videodouyin/videoservice"
	config "mini-douyin/pkg/configs"
	"mini-douyin/pkg/consts"
	"strconv"

	"github.com/cloudwego/kitex/client"
)

var VideoClient videoservice.Client

func initVideo() {
	ip, port, err := config.NacosSelectOneHealthyInstance(consts.VideoServiceName)
	if err != nil {
		panic(err)
	}
	c, err := videoservice.NewClient(consts.VideoServiceName, client.WithHostPorts(ip+":"+strconv.Itoa(int(port))))
	if err != nil {
		panic(err)
	}
	VideoClient = c
	config.Logger.Info("Init Video RPC Client Success!")
}

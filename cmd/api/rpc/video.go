package rpc

import (
	"mini-douyin/cmd/api/kitex_gen/videodouyin/videoservice"
	config "mini-douyin/pkg/configs"
	"mini-douyin/pkg/consts"

	"github.com/cloudwego/kitex/client"
)

var VideoClient videoservice.Client

func initVideo() {
	c, err := videoservice.NewClient(consts.VideoServiceName, client.WithHostPorts("0.0.0.0:9999"))
	if err != nil {
		panic(err)
	}
	VideoClient = c
	config.Logger.Info("Init Video RPC Client Success!")
}

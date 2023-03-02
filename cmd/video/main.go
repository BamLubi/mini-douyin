package main

import (
	"log"
	videodouyin "mini-douyin/cmd/video/kitex_gen/videodouyin/videoservice"
)

func main() {
	svr := videodouyin.NewServer(new(VideoServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}

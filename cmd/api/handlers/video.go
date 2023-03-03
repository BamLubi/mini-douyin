package handlers

import (
	"context"
	"mini-douyin/cmd/api/kitex_gen/videodouyin"
	"mini-douyin/cmd/api/middleware"
	"mini-douyin/cmd/api/rpc"
	config "mini-douyin/pkg/configs"
	"mini-douyin/pkg/consts"
	"mini-douyin/pkg/utils"
	"os"
	"strconv"
	"strings"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func Feed(c *gin.Context) {
	// 提取时间
	latest_time_str := c.Query("latest_time")
	latest_time, _ := strconv.ParseInt(latest_time_str, 10, 64)
	if latest_time == 0 {
		latest_time = time.Now().UnixNano() / int64(time.Millisecond)
	}
	// 提取用户id，可以为空
	var userId int64 = -1
	token_str := c.Query("token")
	if len(token_str) != 0 {
		token, _ := middleware.Jwt.ParseTokenString(token_str)
		claims := jwt.ExtractClaimsFromToken(token)
		userId = int64(claims[consts.IdentityKey].(float64))
	}
	req := new(videodouyin.FeedRequest)
	req.LatestTime = &latest_time
	if userId != -1 {
		req.UserId = &userId
	}
	// RPC获取所有的列表
	resp, err := rpc.VideoClient.Feed(context.Background(), req)
	if err != nil {
		config.Logger.Error(err.Error())
		return
	}
	c.JSON(200, resp)
}

func PublishAction(c *gin.Context) {
	// 提取参数
	token := c.PostForm("token")
	title := c.PostForm("title")
	file, err := c.FormFile("data")
	if err != nil {
		config.Logger.Error(err.Error())
	}

	videoId := utils.IdGen()
	videoIdString := strconv.Itoa(int(videoId))
	videoFileName := videoIdString + "." + strings.Split(file.Filename, ".")[1]
	coverFileName := videoIdString + ".jpg"
	playUrl := "http://192.168.45.129:6789/static/video/" + videoFileName
	coverUrl := "http://192.168.45.129:6789/static/cover/" + coverFileName
	playUrlReal := "public/video/" + videoFileName
	coverUrlReal := "public/cover/" + coverFileName
	userID := middleware.ParseUserIdFromTokenString(token)

	// 协程将文件保存在本地并创建封面图
	go func() {
		if err := c.SaveUploadedFile(file, playUrlReal); err != nil {
			config.Logger.Error(err.Error())
			return
		}
		// 创建首页图片
		if err := utils.GetVideoCover(playUrlReal, coverUrlReal); err != nil {
			config.Logger.Error(err.Error())
			return
		}
		// TODO: 是否需要做失败后的操作，比如将这条数据置为失效
	}()

	req := &videodouyin.PublishActionRequest{UserId: userID, VideoId: videoId, PlayUrl: playUrl, CoverUrl: coverUrl, Title: title}
	resq, err := rpc.VideoClient.PublishAction(context.Background(), req)
	if err != nil {
		config.Logger.Error(err.Error())
		// 删除已经保存的文件
		go func() {
			os.Remove(playUrlReal)
			os.Remove(coverUrlReal)
		}()
		return
	}
	c.JSON(200, resq)
}

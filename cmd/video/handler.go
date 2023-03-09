package main

import (
	"context"
	videodouyin "mini-douyin/cmd/video/kitex_gen/videodouyin"
	config "mini-douyin/pkg/configs"
	"mini-douyin/pkg/entity"
	"strconv"
	"time"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// Feed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Feed(ctx context.Context, req *videodouyin.FeedRequest) (resp *videodouyin.FeedResponse, err error) {
	// 最多 30 个视频，时间倒序，返回最小时间
	// video.create_time < latestTime
	var videoList []*entity.VideoInfo
	err = config.DB.Table("videoinfo").Preload("User").Where("videoinfo.create_time < ?", req.LatestTime).Order("videoinfo.create_time DESC").Limit(30).Find(&videoList).Error
	if err != nil {
		err_msg := "Feed DB error"
		resp = &videodouyin.FeedResponse{StatusCode: 1, StatusMsg: &err_msg}
		return
	}
	videoListIDL := EntityVideoList2IDLVideoList(videoList)
	if len(videoListIDL) == 0 {
		resp = &videodouyin.FeedResponse{StatusCode: 0}
		return
	}
	// 查询redis是否需要更新统计信息
	// 如果redis中有，则添加入返回值
	// 如果没有，则从mysql插入redis
	config.Logger.Debug("运行到这，获取redis")
	for i := 0; i < len(videoListIDL); i++ {
		videoId := strconv.Itoa(int(videoListIDL[i].Id))
		config.Logger.Debug("vide id" + videoId)

		config.RD.Send("SELECT", 10)
		config.RD.Send("HGET", "video_"+videoId, "favorite_count")
		config.RD.Send("HGET", "video_"+videoId, "comment_count")
		config.RD.Flush()
		v, _ := config.RD.Receive()
		v, _ = config.RD.Receive()
		if v == nil {
			config.RD.Send("SELECT", 10)
			config.RD.Send("HSET", "video_"+videoId, "favorite_count", videoListIDL[i].FavoriteCount)
			config.RD.Flush()
		} else {
			str := ""
			for _, n := range v.([]uint8) {
				str += string(n)
			}
			config.Logger.Debug("favorite_count " + str)
			videoListIDL[i].FavoriteCount, _ = strconv.ParseInt(str, 10, 64)
		}
		v, _ = config.RD.Receive()
		if v == nil {
			config.RD.Send("SELECT", 10)
			config.RD.Do("HSET", "video_"+videoId, "comment_count", videoListIDL[i].CommentCount)
			config.RD.Flush()
		} else {
			str := ""
			for _, n := range v.([]uint8) {
				str += string(n)
			}
			config.Logger.Debug("comment_count " + str)
			videoListIDL[i].CommentCount, _ = strconv.ParseInt(str, 10, 64)
		}
	}

	var lastTime int64 = videoListIDL[len(videoListIDL)-1].CreateTime
	resp = &videodouyin.FeedResponse{StatusCode: 0, VideoList: videoListIDL, NextTime: &lastTime}

	return
}

// PublishAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishAction(ctx context.Context, req *videodouyin.PublishActionRequest) (resp *videodouyin.PublishActionResponse, err error) {
	videoInfo := &entity.VideoInfo{
		Id:         req.VideoId,
		UserId:     req.UserId,
		PlayUrl:    req.PlayUrl,
		CoverUrl:   req.CoverUrl,
		Title:      req.Title,
		CreateTime: time.Now().UnixNano() / int64(time.Millisecond),
	}

	err = config.DB.Create(&videoInfo).Error
	if err != nil {
		config.Logger.Error("PublishAction: " + err.Error())

		err_msg := err.Error()
		resp = &videodouyin.PublishActionResponse{StatusCode: 1, StatusMsg: &err_msg}
		return resp, err
	}

	// 返回结果
	resp = &videodouyin.PublishActionResponse{StatusCode: 0}
	return resp, nil
}

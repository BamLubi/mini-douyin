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

func Iface2str(data interface{}) string {
	str := ""
	if v, ok := data.([]uint8); ok {
		for _, n := range v {
			str += string(n)
		}
	} else if v, ok := data.([]int64); ok {
		for _, n := range v {
			str += strconv.Itoa(int(n))
		}
	}
	return str
}

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
	for i := 0; i < len(videoListIDL); i++ {
		videoId := strconv.Itoa(int(videoListIDL[i].Id))
		config.RD.Send("SELECT", 10)
		config.RD.Send("HGET", "video_"+videoId, "favorite_count")
		config.RD.Send("HGET", "video_"+videoId, "comment_count")
		config.RD.Flush()
		v, _ := config.RD.Receive()
		v, _ = config.RD.Receive()
		// 点赞信息是否命中，命中使用缓存，不命中将数据库信息写入redis
		if v == nil {
			go func ()  {
				config.RD.Send("SELECT", 10)
				config.RD.Send("HSET", "video_"+videoId, "favorite_count", videoListIDL[i].FavoriteCount)
				config.RD.Flush()
			}()
		} else {
			videoListIDL[i].FavoriteCount, _ = strconv.ParseInt(Iface2str(v), 10, 64)
		}
		// 评论信息是否命中
		v, _ = config.RD.Receive()
		if v == nil {
			go func ()  {
				config.RD.Send("SELECT", 10)
				config.RD.Do("HSET", "video_"+videoId, "comment_count", videoListIDL[i].CommentCount)
				config.RD.Flush()
			}()
		} else {
			videoListIDL[i].CommentCount, _ = strconv.ParseInt(Iface2str(v), 10, 64)
		}
		// 检查当前用户是否喜欢
		if req.UserId != nil {
			userIdStr := strconv.Itoa(int(*req.UserId))
			config.RD.Send("SELECT", 10)
			config.RD.Send("HGET", "user_favorite_"+userIdStr, videoId)
			config.RD.Flush()
			v, _ = config.RD.Receive()
			v, _ = config.RD.Receive()
			if v == nil {
				// 从数据库中查找，如果点赞则写入redis
				var favorite entity.Favorite
				err = config.DB.Raw("SELECT * FROM favorites WHERE user_id = ? AND video_id = ?", userIdStr, videoId).Scan(&favorite).Error
				if err == nil && favorite.ActionType == 1 {
					videoListIDL[i].IsFavorite = true
					config.RD.Send("SELECT", 10)
					config.RD.Send("HSET", "user_favorite_"+userIdStr, videoId, 1)
					config.RD.Flush()
				}
			} else {
				if Iface2str(v) == "1" {
					videoListIDL[i].IsFavorite = true
				} else {
					videoListIDL[i].IsFavorite = false
				}
			}
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

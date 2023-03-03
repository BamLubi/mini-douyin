package main

import (
	"context"
	"mini-douyin/cmd/video/kitex_gen/base"
	videodouyin "mini-douyin/cmd/video/kitex_gen/videodouyin"
	config "mini-douyin/pkg/configs"
	"mini-douyin/pkg/entity"
	"time"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// Feed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Feed(ctx context.Context, req *videodouyin.FeedRequest) (resp *videodouyin.FeedResponse, err error) {
	// TODO: Your code here...
	var followCount int64 = 0
	var followerCount int64 = 0
	var DemoUser = base.User{
		Id:            1,
		Name:          "TestUser",
		FollowCount:   &followCount,
		FollowerCount: &followerCount,
		IsFollow:      false,
	}
	var DemoVideos = []*base.Video{
		{
			Id:            1,
			Author:        &DemoUser,
			PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
			CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
			FavoriteCount: 0,
			CommentCount:  0,
			IsFavorite:    false,
		},
	}
	resp = &videodouyin.FeedResponse{StatusCode: 0, VideoList: DemoVideos, NextTime: req.LatestTime}
	return
}

// PublishAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishAction(ctx context.Context, req *videodouyin.PublishActionRequest) (resp *videodouyin.PublishActionResponse, err error) {
	videoInfo := &entity.VideoInfo{
		Id:       req.VideoId,
		UserId:   req.UserId,
		PlayUrl:  req.PlayUrl,
		CoverUrl: req.CoverUrl,
		Title:    req.Title,
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

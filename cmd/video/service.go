package main

import (
	"mini-douyin/cmd/video/kitex_gen/base"
	"mini-douyin/pkg/entity"
)

func EntityUserInfo2IDLUser(userinfo *entity.UserInfo) (user *base.User) {
	user = &base.User{
		Id:              userinfo.Id,
		Name:            userinfo.Name,
		FollowCount:     &userinfo.FollowCount,
		FollowerCount:   &userinfo.FollowerCount,
		IsFollow:        userinfo.IsFollow,
		Avatar:          &userinfo.Avatar,
		BackgroundImage: &userinfo.BackgroundImage,
		Signature:       &userinfo.Signature,
		TotalFavorited:  &userinfo.TotalFavorited,
		WorkCount:       &userinfo.WorkCount,
		FavoriteCount:   &userinfo.FavoriteCount,
	}
	return
}

func EntityVideoList2IDLVideoList(videoList []*entity.VideoInfo) []*base.Video {
	var videoListIDL []*base.Video
	for _, video := range videoList {
		videoIDL := &base.Video{
			Id:            video.Id,
			Author:        EntityUserInfo2IDLUser(&video.User),
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    video.IsFavorite,
			Title:         video.Title,
			CreateTime:    video.CreateTime,
		}

		videoListIDL = append(videoListIDL, videoIDL)
	}
	return videoListIDL
}

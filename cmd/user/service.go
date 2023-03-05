package main

import (
	"mini-douyin/cmd/user/kitex_gen/base"
	config "mini-douyin/pkg/configs"
	"mini-douyin/pkg/entity"
)

// 注册事务
func RegisterTransaction(user *entity.User, userinfo *entity.UserInfo) (err error) {
	tx := config.DB.Begin()
	// 插入用户表，记录用户名密码等
	err = tx.Create(&user).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	// 插入用户信息表，记录用户基本信息
	err = tx.Create(&userinfo).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

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

func EntityFavoriteList2IDLVideoList(favoriteList []*entity.Favorite) []*base.Video {
	var videoListIDL []*base.Video
	for _, favorite := range favoriteList {
		video := &favorite.Video
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

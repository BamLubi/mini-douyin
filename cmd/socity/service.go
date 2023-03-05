package main

import (
	"mini-douyin/cmd/socity/kitex_gen/base"
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

func EntityComment2IDLComment(comment *entity.Comment, user *base.User) (commenIDL *base.Comment) {
	commenIDL = &base.Comment{
		Id:         comment.Id,
		User:       user,
		Content:    comment.Content,
		CreateDate: comment.CreateDate,
	}
	return
}

func EntityCommentList2IDLCommentList(commentList []*entity.Comment) []*base.Comment{
	var commentListIDL []*base.Comment
	for _, comment := range commentList {
		commentILD := &base.Comment{
			Id: comment.Id,
			User: EntityUserInfo2IDLUser(&comment.User),
			Content: comment.Content,
			CreateDate: comment.CreateDate,
		}

		commentListIDL = append(commentListIDL, commentILD)
	}
	return commentListIDL
}

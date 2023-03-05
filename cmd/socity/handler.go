package main

import (
	"context"
	socitydouyin "mini-douyin/cmd/socity/kitex_gen/socitydouyin"
	"mini-douyin/cmd/socity/kitex_gen/userdouyin"
	"mini-douyin/cmd/socity/rpc"
	config "mini-douyin/pkg/configs"
	"mini-douyin/pkg/entity"
	"mini-douyin/pkg/utils"
	"time"
)

// SocityServiceImpl implements the last service interface defined in the IDL.
type SocityServiceImpl struct{}

// FavoriteAction implements the SocityServiceImpl interface.
func (s *SocityServiceImpl) FavoriteAction(ctx context.Context, req *socitydouyin.FavoriteActionRequest) (resp *socitydouyin.FavoriteActionResponse, err error) {
	// 无论是点赞还是取消点赞都修改数据库
	// TODO：使用redis优化时在考虑
	if req.ActionType == 1 {
		favourite := &entity.Favorite{Id: utils.IdGen(), UserId: req.UserId, VideoId: req.VideoId, ActionType: req.ActionType}
		err = config.DB.Create(favourite).Error
	} else {
		err = config.DB.Model(&entity.Favorite{}).Where("user_id = ?", req.UserId).Where("video_id = ?", req.VideoId).Update("action_type", req.ActionType).Error
	}
	if err != nil {
		config.Logger.Error(err.Error())
		err_msg := "FavoriteAction DB error"
		resp = &socitydouyin.FavoriteActionResponse{StatusCode: 1, StatusMsg: &err_msg}
		return
	}
	resp = &socitydouyin.FavoriteActionResponse{StatusCode: 0}
	return
}

// CommentAction implements the SocityServiceImpl interface.
func (s *SocityServiceImpl) CommentAction(ctx context.Context, req *socitydouyin.CommentActionRequest) (resp *socitydouyin.CommentActionResponse, err error) {
	if req.ActionType == 1 {
		nowDate := time.Now().Format("01-02")
		comment := &entity.Comment{Id: utils.IdGen(), UserId: req.UserId, VideoId: req.VideoId, Content: *req.CommentText, CreateDate: nowDate, Active: 1}
		// err = config.DB.Create(comment).Error
		err = config.DB.Save(comment).Error
		if err != nil {
			config.Logger.Error(err.Error())
			err_msg := "FavoriteAction DB error"
			resp = &socitydouyin.CommentActionResponse{StatusCode: 1, StatusMsg: &err_msg}
			return
		}
		// 调用rpc获取当前用户信息
		userInfoReq := &userdouyin.UserInfoRequest{UserId: req.UserId}
		var userInfoResq *userdouyin.UserInfoResponse
		userInfoResq, err = rpc.UserClient.UserInfo(ctx, userInfoReq)
		if err != nil {
			err_msg := "FavoriteAction DB error"
			resp = &socitydouyin.CommentActionResponse{StatusCode: 1, StatusMsg: &err_msg}
			return
		} // TODO: 出错需要回滚
		user := userInfoResq.User
		commentIDL := EntityComment2IDLComment(comment, user)
		resp = &socitydouyin.CommentActionResponse{StatusCode: 0, Comment: commentIDL}
		return
	} else {
		err = config.DB.Model(&entity.Comment{}).Where("id = ?", req.CommentId).Update("active", 0).Error
		if err != nil {
			config.Logger.Error(err.Error())
			err_msg := "FavoriteAction DB error"
			resp = &socitydouyin.CommentActionResponse{StatusCode: 1, StatusMsg: &err_msg}
			return
		}
		resp = &socitydouyin.CommentActionResponse{StatusCode: 0}
		return
	}
}

// CommentList implements the SocityServiceImpl interface.
func (s *SocityServiceImpl) CommentList(ctx context.Context, req *socitydouyin.CommentListRequest) (resp *socitydouyin.CommentListResponse, err error) {
	var commentList []*entity.Comment
	err = config.DB.Table("comment").Preload("User").Where("comment.video_id = ?", req.VideoId).Order("comment.create_date DESC").Find(&commentList).Error
	if err != nil {
		config.Logger.Error(err.Error())
		err_msg := "comment DB error"
		resp = &socitydouyin.CommentListResponse{StatusCode: 1, StatusMsg: &err_msg}
		return
	}
	commentListIDL := EntityCommentList2IDLCommentList(commentList)
	resp = &socitydouyin.CommentListResponse{StatusCode: 0, CommentList: commentListIDL}

	return
}

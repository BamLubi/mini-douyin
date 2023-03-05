package handlers

import (
	"context"
	"mini-douyin/cmd/api/kitex_gen/socitydouyin"
	"mini-douyin/cmd/api/middleware"
	"mini-douyin/cmd/api/rpc"
	config "mini-douyin/pkg/configs"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FavoriteAction(c *gin.Context) {
	userId := middleware.ParseUserIdFromTokenString(c.Query("token"))
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 32)
	req := &socitydouyin.FavoriteActionRequest{UserId: userId, VideoId: videoId, ActionType: int32(actionType)}
	resp, err := rpc.SocityClient.FavoriteAction(context.Background(), req)
	if err != nil {
		config.Logger.Error(err.Error())
	}
	c.JSON(200, resp)
}

func CommentAction(c *gin.Context) {
	userId := middleware.ParseUserIdFromTokenString(c.Query("token"))
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 32)
	// 1->comment, 2->decomment
	var req *socitydouyin.CommentActionRequest
	if int32(actionType) == 1 {
		commentText := c.Query("comment_text")
		req = &socitydouyin.CommentActionRequest{UserId: userId, VideoId: videoId, ActionType: int32(actionType), CommentText: &commentText}
	} else {
		commentId, _ := strconv.ParseInt(c.Query("comment_id"), 10, 64)
		req = &socitydouyin.CommentActionRequest{UserId: userId, VideoId: videoId, ActionType: int32(actionType), CommentId: &commentId}
	}
	resp, err := rpc.SocityClient.CommentAction(context.Background(), req)
	if err != nil {
		config.Logger.Error(err.Error())
	}
	c.JSON(200, resp)
}

func CommentList(c *gin.Context) {
	userId := middleware.ParseUserIdFromTokenString(c.Query("token"))
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	req := &socitydouyin.CommentListRequest{UserId: userId, VideoId: videoId}
	resp, err := rpc.SocityClient.CommentList(context.Background(), req)
	if err != nil {
		config.Logger.Error(err.Error())
	}
	c.JSON(200, resp)
}

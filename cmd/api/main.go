package main

import (
	"mini-douyin/cmd/api/handlers"
	"mini-douyin/cmd/api/middleware"
	"mini-douyin/cmd/api/rpc"
	config "mini-douyin/pkg/configs"

	"github.com/gin-gonic/gin"
)

func Init() {
	config.InitConfigs("/root/mini-douyin/logs/api.log", "debug")
	rpc.InitRPC()
	middleware.InitJWT()
}

func main() {
	Init()

	r := gin.Default()

	// 解决跨域
	// r.Use(web.CorsConfig())

	// 初始化路由
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")
	{
		apiRouter.POST("/user/register/", handlers.UserRegister)                                      // 注册接口 - user
		apiRouter.POST("/user/login/", middleware.Jwt.LoginHandler)                                   // 登录接口 - user
		apiRouter.GET("/user/", middleware.Jwt.MiddlewareFunc(), handlers.UserInfo)                   // 用户信息 - user
		apiRouter.GET("/publish/list/", middleware.Jwt.MiddlewareFunc(), handlers.PublishList)        // 发布列表 - user
		apiRouter.POST("/publish/action/", middleware.Jwt.MiddlewareFunc(), handlers.PublishAction)   // 视频投稿 - video
		apiRouter.GET("/feed/", handlers.Feed)                                                        // 视频流  - video
		apiRouter.POST("/favorite/action/", middleware.Jwt.MiddlewareFunc(), handlers.FavoriteAction) // 赞操作 - socity
		apiRouter.POST("/comment/action/", middleware.Jwt.MiddlewareFunc(), handlers.CommentAction)   // 评论操作 - socity
		apiRouter.GET("/comment/list/", middleware.Jwt.MiddlewareFunc(), handlers.CommentList)        // 视频评论列表 - socity
		apiRouter.GET("/favorite/list/", middleware.Jwt.MiddlewareFunc(), handlers.FavoriteList)      // 喜欢列表 - user
	}

	// // extra apis - II
	// apiRouter.POST("/relation/action/", controller.RelationAction) // 关系操作
	// apiRouter.GET("/relation/follow/list/", controller.FollowList) // 用户关注列表 - user
	// apiRouter.GET("/relation/follower/list/", controller.FollowerList) // 用户粉丝列表 - user
	// apiRouter.GET("/relation/friend/list/", controller.FriendList) // 用户好友列表 - user
	// apiRouter.GET("/message/chat/", controller.MessageChat) // 聊天记录
	// apiRouter.POST("/message/action/", controller.MessageAction) // 消息操作

	r.Run("0.0.0.0:6789")
}

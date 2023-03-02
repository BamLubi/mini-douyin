package handlers

import (
	"context"
	"mini-douyin/cmd/api/kitex_gen/userdouyin"
	"mini-douyin/cmd/api/middleware"
	"mini-douyin/cmd/api/rpc"

	config "mini-douyin/pkg/configs"

	"github.com/gin-gonic/gin"
)

// 用户注册
func UserRegister(c *gin.Context) {
	// var req userdouyin.UserRegisterRequest
	// if err := c.Bind(&req); err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// 获取参数，参数为 Param 格式
	req := &userdouyin.UserRegisterRequest{Username: c.Query("username"), Password: c.Query("password")}
	// 调用RPC执行用户注册，返回用户的id即可
	resp, err := rpc.UserClient.UserRegister(context.Background(), req)
	if err != nil {
		config.Logger.Error(err.Error())
		return
	}
	// 根据用户的id生成相应的token
	if resp.StatusCode == 0 {
		token, _, err := middleware.Jwt.TokenGenerator(resp.UserId)
		if err != nil {
			config.Logger.Error(err.Error())
			return
		}
		resp.Token = token
	}
	c.JSON(200, resp)
}

func UserInfo(c *gin.Context) {

}

func PublishList(c *gin.Context) {

}

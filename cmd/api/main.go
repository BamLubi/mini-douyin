package main

import (
	"mini-douyin/cmd/api/web"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 解决跨域
	r.Use(web.CorsConfig())

	// 初始化路由
	web.InitRouter(r)

	r.Run()
}

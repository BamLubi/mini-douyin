package handlers

import (
	"log"
	"mini-douyin/cmd/api/kitex_gen/videodouyin"
	"mini-douyin/cmd/api/middleware"
	"mini-douyin/pkg/consts"
	"strconv"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// /douyin/feed?latest_time=1677758442403&token=eyJhbGci...

func Feed(c *gin.Context) {
	// 提取时间
	latest_time_str := c.Query("latest_time")
	latest_time, _ := strconv.ParseInt(latest_time_str, 10, 64)
	if latest_time == 0 {
		latest_time = time.Now().UnixNano() / int64(time.Millisecond)
	}
	// 提取用户id，可以为空
	var userId int64 = -1
	token_str := c.Query("token")
	if len(token_str) != 0 {
		token, _ := middleware.Jwt.ParseTokenString(token_str)
		claims := jwt.ExtractClaimsFromToken(token)
		userId = int64(claims[consts.IdentityKey].(float64))
	}
	req := new(videodouyin.FeedRequest)
	req.LatestTime = &latest_time
	if userId != -1 {
		// req.Token
	}

	log.Println(latest_time)
}

func Publish(c *gin.Context) {

}

package middleware

import (
	"errors"
	"log"
	"mini-douyin/cmd/api/kitex_gen/userdouyin"
	"mini-douyin/cmd/api/rpc"
	"mini-douyin/pkg/consts"
	"time"

	config "mini-douyin/pkg/configs"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var Jwt *jwt.GinJWTMiddleware

func InitJWT() {
	JwtMiddleware, _ := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte(consts.SecretKey),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: consts.IdentityKey,
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return int64(claims[consts.IdentityKey].(float64))
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			// var req userdouyin.UserLoginRequest
			// if err := c.ShouldBind(&req); err != nil {
			// 	return "", jwt.ErrMissingLoginValues
			// }
			// 获取参数，参数为 Param 格式
			req := &userdouyin.UserLoginRequest{Username: c.Query("username"), Password: c.Query("password")}
			if len(req.Username) == 0 || len(req.Password) == 0 {
				return "", jwt.ErrMissingLoginValues
			}
			resp, err := rpc.UserClient.UserLogin(c, req)
			if err != nil {
				return -1, err
			}
			if resp.StatusCode != 0 {
				return -1, errors.New("password wrong")
			}
			return resp.UserId, nil
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{consts.IdentityKey: v}
			}
			return jwt.MapClaims{}
		},
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			resp := new(userdouyin.UserLoginResponse)
			resp.StatusCode = 0
			resp.Token = token
			resp.UserId = ParseUserIdFromTokenString(token)
			log.Println("Login:", resp.UserId, resp.Token)
			c.JSON(200, resp)
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenHeadName: "Bearer",
		TokenLookup:   "header: Authorization, query: token, cookie: jwt, param: token, form: token",
		TimeFunc:      time.Now,
	})

	Jwt = JwtMiddleware
}

func ParseUserIdFromTokenString(token_str string) int64 {
	token, _ := Jwt.ParseTokenString(token_str)
	claims := jwt.ExtractClaimsFromToken(token)
	userId := int64(claims[consts.IdentityKey].(float64))
	return userId
}

func GenTokenStringFromUserId(userId int64) (string, error) {
	token, _, err := Jwt.TokenGenerator(userId)
	if err != nil {
		config.Logger.Error(err.Error())
		return "", err
	}
	return token, nil
}

func CheckUserId2TokenString(userId int64, token_str string) error {
	if ParseUserIdFromTokenString(token_str) != userId {
		return errors.New("id is not match token")
	}
	return nil
}

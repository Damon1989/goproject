package middleware

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/damon/gogofly/api"
	"github.com/damon/gogofly/global"
	"github.com/damon/gogofly/global/constants"
	"github.com/damon/gogofly/model"
	"github.com/damon/gogofly/service"
	"github.com/damon/gogofly/utils"
	"github.com/gin-gonic/gin"
)

const (
	ERR_CODE_INVALID_TOKEN     = 10401 // 无效的token
	ERR_CODE_TOKEN_PARSED      = 10402 // token解析错误
	ERR_CODE_TOKEN_NOT_MATCHED = 10403 // token不匹配
	ERR_CODE_TOKEN_EXPIRED     = 10404 // token已过期
	ERR_CODE_TOKEN_RENEWED     = 10405 // token续期失败
	TOKEN_NAME                 = "Authorization"
	TOKEN_PREFIX               = "Bearer: "
	RENEW_TOKEN_DURATION       = 10 * 60 * time.Second
)

func tokenErr(c *gin.Context, code int) {
	api.Fail(c, api.ResponseJson{
		Status: http.StatusUnauthorized,
		Code:   code,
		Msg:    "invalid token",
	})
	c.Abort()
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 token
		token := c.GetHeader(TOKEN_NAME)

		if token == "" || !strings.HasPrefix(token, TOKEN_PREFIX) {
			tokenErr(c, ERR_CODE_INVALID_TOKEN)
			return
		}

		token = token[len(TOKEN_PREFIX):]

		iJwtCustomClaims, err := utils.ParseToken(token)
		nUserId := iJwtCustomClaims.ID
		if err != nil {
			tokenErr(c, ERR_CODE_TOKEN_PARSED)
			return
		}
		stUserId := strconv.Itoa(int(nUserId))
		stRedisUserIDKey := strings.Replace(constants.LOGIN_USER_TOKEN_REDIS_KEY, "{ID}", stUserId, -1)
		// Token与访问者登录对应的token不一致，直接返回

		stRedisToken, err := global.RedisClient.Get(stRedisUserIDKey)
		if err != nil || stRedisToken != token {
			tokenErr(c, ERR_CODE_TOKEN_NOT_MATCHED)
			return
		}

		// Token已过期，直接返回
		nTokenExpireDuration, err := global.RedisClient.GetExpireDuration(stRedisUserIDKey)
		if err != nil || nTokenExpireDuration <= 0 {
			tokenErr(c, ERR_CODE_TOKEN_EXPIRED)
			return
		}

		// Token续期
		if nTokenExpireDuration.Seconds() < RENEW_TOKEN_DURATION.Seconds() {
			stNewToken, err := service.GenerateAndCacheLoginUserToken(nUserId, iJwtCustomClaims.Name)
			if err != nil {
				tokenErr(c, ERR_CODE_TOKEN_RENEWED)
				return
			}
			c.Header(TOKEN_NAME, TOKEN_PREFIX+stNewToken)
		}
		/*	iUser, err := dao.NewUserDao().GetUserById(nUserId)
			if err != nil || iUser.ID == 0 {
				tokenErr(c)
				return
			}
			c.Set(constants.LOGIN_USER, iUser)*/
		c.Set(constants.LOGIN_USER, model.LoginUser{
			ID:   nUserId,
			Name: iJwtCustomClaims.Name,
		})
		c.Next()
	}
}

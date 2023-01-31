package middleware

import (
	"gin_shop/pkg/e"
	"gin_shop/pkg/util"
	"github.com/gin-gonic/gin"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		var code int
		code = 200
		token := context.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ErrorAuthToken
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ErrorAuthCheckToken
			}
		}
		if code != e.Success {
			context.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
			})
			context.Abort()
			return
		}
		context.Next()
	}
}

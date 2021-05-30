package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"service/pkg/app"
	"service/pkg/errcode"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			ecode = errcode.Success
		)
		// 从请求中获取token
		if s, exist := c.GetQuery("token"); exist {
			token = s
		} else {
			token = c.GetHeader("token")
		}
		// 判断token的有效性
		if token == "" {
			ecode = errcode.InvalidParams.WithDetails("未提供有效token")
		} else {
			_, err := app.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					ecode = errcode.UnauthorizedTokenTimeout
				default:
					ecode = errcode.UnauthorizedTokenError
				}
			}
		}
		// 错误码处理
		if ecode != errcode.Success {
			response := app.NewResponse(c)
			response.ToErrResponse(ecode)
			c.Abort()
			return
		}
		c.Next()
	}
}

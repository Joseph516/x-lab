package api

import (
	"github.com/gin-gonic/gin"
	"service/internal/service"
	"service/pkg/app"
	"service/pkg/errcode"
)

type Register struct {

}

func NewRegister() Register {
	return Register{}
}

// Register 注册新用户
func (r Register) Register(c *gin.Context)  {
	param := service.RegisterRequest{}
	response := app.NewResponse(c)

	// 参数绑定与校验
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		response.ToErrResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	// 两次密码判定
	if param.Password != param.Repassword {
		response.ToErrResponse(errcode.PasswordNotSame)
		return
	}
	svc := service.New(c.Request.Context())
	// 将用户信息存入数据库
	err := svc.Register(&param)
	if err != nil {
		response.ToErrResponse(errcode.RegisterFailed.WithDetails(err.Error()))
		return
	}
	response.ToErrResponse(errcode.Success)
}
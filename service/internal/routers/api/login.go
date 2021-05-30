package api

import (
	"github.com/gin-gonic/gin"
	"service/internal/service"
	"service/pkg/app"
	"service/pkg/errcode"
)

type Login struct {
}

func NewLogin() Login {
	return Login{}
}

// LoginIn 用户登陆
func (l Login) LoginIn(c *gin.Context) {
	param := service.LoginRequest{}
	response := app.NewResponse(c)

	// 参数绑定与校验
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		response.ToErrResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 用户校验
	svc := service.New(c.Request.Context())
	err := svc.LoginIn(&param)
	if err != nil {
		response.ToErrResponse(errcode.LoginInFailed.WithDetails(err.Error()))
		return
	}
	// 登陆成功，生成token
	token, err := app.GenerateToken(param.Username, param.Password)
	if err != nil {
		response.ToErrResponse(errcode.ServerError)
		return
	}
	data := gin.H{
		"code": errcode.Success.Code(),
		"msg":  errcode.Success.Msg(),
		"data": gin.H{
			"token": token,
		},
	}
	response.ToResponse(data)
}

func (l Login) Exit(c *gin.Context) {
	// TBD
}

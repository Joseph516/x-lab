/*
package app 是对返回响应进行封装
*/

package app

import (
	"github.com/gin-gonic/gin"
	"x-lab/service/pkg/errcode"
)

type Response struct {
	Ctx *gin.Context
}

func NewResponse(c *gin.Context) *Response {
	return &Response{Ctx: c}
}

// ToResponse 请求成功，返回响应的数据
func (resp *Response) ToResponse(data interface{}) {
	if data == nil {
		resp.Ctx.JSON(errcode.Success.ToHttpStatusCode(), gin.H{})
		return
	}
	resp.Ctx.JSON(errcode.Success.ToHttpStatusCode(), data)
}

// ToErrResponse 请求失败，返回错误信息
func (resp *Response) ToErrResponse(err *errcode.Error) {
	response := gin.H{
		"code": err.Code(), // err.ToHttpStatusCode()
		"msg":  err.Msg(),
	}
	if len(err.Details()) > 0 {
		response["details"] = err.Details()
	}
	resp.Ctx.JSON(err.ToHttpStatusCode(), response)
}

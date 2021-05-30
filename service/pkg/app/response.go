/*
package app 是对返回响应进行封装
*/

package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"service/pkg/errcode"
)

type Response struct {
	Ctx *gin.Context
}

func NewResponse(c *gin.Context) *Response {
	return &Response{Ctx: c}
}

// ToResponse 请求成功，返回响应的数据
func (resp *Response) ToResponse(data interface{}) {
	response := gin.H{
		"code": errcode.Success.Code(),
		"msg":  errcode.Success.Msg(),
	}
	if data != nil {
		response["data"] = data
	}
	resp.Ctx.JSON(http.StatusOK, response)
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
	resp.Ctx.JSON(http.StatusOK, response)
}

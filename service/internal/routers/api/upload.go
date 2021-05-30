/*
package api upload 管理文件的上传功能
*/

package api

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"service/internal/service"
	"service/pkg/app"
	"service/pkg/errcode"
	"service/pkg/upload"
	"strconv"
)

type Upload struct {
}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) CheckFile(c *gin.Context) (multipart.File, *multipart.FileHeader, int) {
	file, fileHeader, err := c.Request.FormFile("file")
	response := app.NewResponse(c)
	if err != nil {
		response.ToErrResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return nil, nil, 0
	}
	fileType, _ := strconv.Atoi(c.PostForm("type"))
	if fileHeader == nil || fileType <= 0 {
		response.ToErrResponse(errcode.InvalidParams.WithDetails("Invalid file type."))
		return nil, nil, 0
	}
	return file, fileHeader, fileType
}

func (u Upload) UploadFile(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")
	response := app.NewResponse(c)
	if err != nil {
		response.ToErrResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}
	fileType, _ := strconv.Atoi(c.PostForm("type"))
	if fileHeader == nil || fileType <= 0 {
		response.ToErrResponse(errcode.InvalidParams.WithDetails("Invalid file type."))
		return
	}
	// 从token中获取当前用户
	token := c.GetHeader("token")
	claims, err := app.ParseToken(token)
	if err != nil {
		response.ToErrResponse(errcode.UnauthorizedTokenError.WithDetails(err.Error()))
		return
	}
	// 将文件保存到指定位置
	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader, claims.AppKey)
	if err != nil {
		response.ToErrResponse(errcode.UploadFileFailed.WithDetails(err.Error()))
		return
	}
	// 上传成功
	data := gin.H{
		"file_access_url": fileInfo.AccessUrl,
	}
	response.ToResponse(data)
}

func (u Upload) Get(c *gin.Context) {

}

func (u Upload) List(c *gin.Context) {

}

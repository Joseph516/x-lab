package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"service/global"
	"service/internal/routers/api"
)

func NewRouters() *gin.Engine {
	r := gin.New()

	// 文件上传
	upload := api.NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	// 提供静态资源访问，当访问/static，实际上访问的是global.AppSetting.UploadSavePath文件下的文件
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	// v1 := r.Group("/v1")
	// {
	// 	v1.POST("/login", loginEndpoint)
	// }

	return r
}
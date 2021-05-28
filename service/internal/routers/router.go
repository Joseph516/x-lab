package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"service/global"
	"service/internal/middleware"
	"service/internal/routers/api"
)

func NewRouters() *gin.Engine {
	r := gin.New()

	// 中间件设置
	r.Use(middleware.Cors())

	// 注册与登陆
	register := api.NewRegister()
	login := api.NewLogin()
	user := r.Group("/user")
	{
		user.POST("/register", register.Register)
		user.POST("/login", login.LoginIn)
	}

	// 文件上传
	upload := api.NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	// 提供静态资源访问，当访问/static，实际上访问的是global.AppSetting.UploadSavePath文件下的文件
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	r.GET("/upload/file", upload.List)
	r.GET("/upload/file:filename", upload.Get)



	return r
}

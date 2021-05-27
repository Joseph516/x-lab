package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"service/global"
	"service/internal/model"
	"service/internal/routers"
	"service/pkg/settings"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("setupSetting failed: %v", err)
	}

	err = dbConnectSetting()
	if err != nil {
		log.Fatalf("配置数据库连接失败：%v", err)
	}
}

func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouters()
	s := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("Starting server failed: %v", err)
	}
}

/***************************/
/********** 配置 ************/
/***************************/
// 配置热更新
func setupSetting() error {
	// 服务器基本配置
	setting, err := settings.NewSetting("configs/")
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.WriteTimeout *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second

	// 数据库配置
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	// App配置
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	return nil
}

// 配置数据库连接
func dbConnectSetting() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

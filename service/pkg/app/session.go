package app

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"service/global"
)

func SetSession(c *gin.Context, encryptedKey, username string) error {
	// 设置session后会将用户登陆状态设置到cookies，后浏览器进行请求时会自动带上cookies。
	session := sessions.Default(c)
	session.Set(encryptedKey, username)
	return session.Save()
}

func DeleteSession(c *gin.Context, encryptedKey string) error {
	// 退出时清除sessions
	session := sessions.Default(c)
	session.Delete(global.JWTSetting.SessionSecretLogin)
	// session.Clear() // 清除所有session
	return session.Save()
}

func GetUserFromSession(c *gin.Context, encryptedKey string) (bool, string) {
	session := sessions.Default(c)
	data := session.Get(encryptedKey)
	if data == nil {
		return false, ""
	}
	username := fmt.Sprintf("%v", data)
	return true, username
}

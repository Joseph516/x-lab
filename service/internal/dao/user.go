package dao

import (
	"errors"
	"service/internal/model"
	"service/pkg/util"
	"time"
)

// Register 用户注册
func (d *Dao) Register(name string, password string) error {
	if ok := d.UserExist(name); ok {
		return errors.New("用户已经存在，请直接登陆")
	}

	user := model.User{
		Username:   name,
		Password:   util.EncodeMD5(password),
		Status:     0,
		CreateTime: time.Now().Unix(),
	}
	return user.Create(d.engine)
}

// UserExist 根据name判断数据库中用户是否存在
func (d *Dao) UserExist(name string) bool {
	user := model.User{
		Username:   name,
	}
	users, _ := user.GetUserListByName(d.engine)
	if len(users) != 0 {
		return true
	}
	return false
}

// 用户登陆
func (d *Dao) LoginIn(name string, password string) error  {
	user := model.User{
		Username:   name,
		Password:   util.EncodeMD5(password),
	}
	_, err := user.GetUser(d.engine)
	return err
}

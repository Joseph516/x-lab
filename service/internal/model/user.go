package model

import "github.com/jinzhu/gorm"

type User struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Status     int    `json:"status"`
	CreateTime int64  `json:"create_time"`
}

// TableName 不能省略，让gorm内部调用，实现在指定数据库表上操作
func (u User) TableName() string {
	return "users"
}

func (u User) GetUser(db *gorm.DB) (*User, error) {
	var user User
	err := db.Where("username = ? AND password = ?", u.Username, u.Password).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u User) GetUserListByName(db *gorm.DB) ([]*User, error) {
	var users []*User
	err := db.Where("username = ?", u.Username).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Create 根据User对象新建数据表数据
func (u User) Create(db *gorm.DB) error {
	return db.Create(&u).Error
}

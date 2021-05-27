package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"service/pkg/settings"
)
import "fmt"

func NewDBEngine(databaseSetting *settings.DatabaseSetting) (*gorm.DB, error) {
	// user:password@/dbname?charset=utf8&parseTime=True&loc=Local
	sqlCom := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	)
	// gorm只能手动创建数据库
	// databaseSetting.DBType: mysql
	db, err := gorm.Open(databaseSetting.DBType, sqlCom)
	if err != nil {
		return nil, err
	}
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)

	return db, nil
}

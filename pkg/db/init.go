package db

import (
	"douyin/pkg/conf"
	"douyin/pkg/repository"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() error {
	dbConnect := conf.DB
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", dbConnect.Username, dbConnect.Password, dbConnect.Host, dbConnect.Port, dbConnect.DbName, dbConnect.Timeout)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DB.AutoMigrate(&repository.User{}, &repository.Video{})
	return nil
}

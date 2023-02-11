package db

import (
	"douyin/pkg/conf"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() error {
	dbConnect := conf.DB
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", dbConnect.Username, dbConnect.Password, dbConnect.Host, dbConnect.Port, dbConnect.DbName, dbConnect.Timeout)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt:            true, //预编译语句
		SkipDefaultTransaction: true, //跳过默认事务
	})
	if err != nil {
		return err
	}
	//err = DB.AutoMigrate(&repository.User{}, &repository.Video{})
	//if err != nil {
	//	return err
	//}
	return nil
}

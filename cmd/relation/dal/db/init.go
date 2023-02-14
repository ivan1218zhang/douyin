package db

import (
	"douyin/pkg/constants"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
)

var DB *gorm.DB

// Init init DB
func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(constants.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)

	if err != nil {
		panic(err)
	}

	//DB.AutoMigrate(&repository.User{}, &repository.FollowedUser{}, &repository.FollowerUser{})
	if err = DB.Use(gormopentracing.New()); err != nil {
		panic(err)
	}

}

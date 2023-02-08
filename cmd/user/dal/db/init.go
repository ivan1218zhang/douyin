package db

import (
	"douyin/pkg/constants"
	"douyin/pkg/repository"
	"douyin/pkg/util"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
)

var DB *gorm.DB
var snowflake *util.SnowFlake
var machineID = 1

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
	DB.AutoMigrate(&repository.User{})
	if err = DB.Use(gormopentracing.New()); err != nil {
		panic(err)
	}

	snowflake = util.NewSnowFlake(int64(machineID))
}

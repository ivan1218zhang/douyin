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

// Init init DB
func Init() {
	snowflake = util.NewSnowFlake(constants.CommentServiceMachineID)

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

	m := DB.Migrator()
	if m.HasTable(&repository.Comment{}) {
		return
	}

	// 表不存在则在数据库中创建表
	err = m.DropTable(&repository.Comment{})
	if err != nil {
		panic(err)
	}
	if err = m.AutoMigrate(&repository.Comment{}); err != nil {
		panic(err)
	}

	if err = DB.Use(gormopentracing.New()); err != nil {
		panic(err)
	}

}

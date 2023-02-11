package repository

import (
	"douyin/pkg/constants"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	UserName  string
	Password  string
}

func (u *User) TableName() string {
	return constants.UserTableName
}

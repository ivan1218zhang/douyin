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
	DeletedAt gorm.DeletedAt
	UserName  string `gorm:"type:varchar(20);Index:idx_user_username"`
	Password  string `gorm:"type:varchar(20)"`
}

func (u *User) TableName() string {
	return constants.UserTableName
}

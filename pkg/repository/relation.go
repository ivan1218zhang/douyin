package repository

import (
	"douyin/pkg/constants"
	"gorm.io/gorm"
	"time"
)

type Relation struct {
	Id         int64
	UserId     int64
	FollowerId int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (u *Relation) TableName() string {
	return constants.RelationTableName
}

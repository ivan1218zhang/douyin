package repository

import (
	"douyin/pkg/constants"
	"gorm.io/gorm"
	"time"
)

type Relation struct {
	Id         int64
	UserId     int64 `gorm:"Index:idx_relation,priority:11"`
	FollowerId int64 `gorm:"Index:idx_relation,priority:12;Index:idx_relation_follower;"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}

func (u *Relation) TableName() string {
	return constants.RelationTableName
}

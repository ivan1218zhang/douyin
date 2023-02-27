package repository

import (
	"douyin/pkg/constants"
	"time"

	"gorm.io/plugin/soft_delete"
)

type Relation struct {
	Id         int64
	UserId     int64 `gorm:"Index:idx_relation,priority:11;uniqueIndex:relation"`
	FollowerId int64 `gorm:"Index:idx_relation,priority:12;Index:idx_relation_follower;uniqueIndex:relation"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  soft_delete.DeletedAt `gorm:"uniqueIndex:relation;softDelete:milli"`
}

func (u *Relation) TableName() string {
	return constants.RelationTableName
}

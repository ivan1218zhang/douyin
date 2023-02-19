package repository

import "douyin/pkg/constants"

type Relation struct {
	Id         int64
	UserId     int64
	FollowerId int64
}

func (u *Relation) TableName() string {
	return constants.RelationTableName
}

package repository

import (
	"gorm.io/gorm"
	"time"
)

type FollowedUser struct {
	ID             int64
	UserId         int64
	FollowedUserId int64
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
}

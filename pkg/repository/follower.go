package repository

import (
	"gorm.io/gorm"
	"time"
)

type FollowerUser struct {
	ID             int64
	UserId         int64
	FollowerUserId int64
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
}

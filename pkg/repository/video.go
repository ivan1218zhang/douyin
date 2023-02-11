package repository

import (
	"douyin/pkg/constants"
	"encoding"
	"encoding/json"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type Video struct {
	ID            int64
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	Title         string
	AuthorId      int64
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64
	CommentCount  int64
}

func (v *Video) TableName() string {
	return constants.VideoTableName
}

// MarshalBinary implements encoding.BinaryMarshaler for redis
func (v *Video) MarshalBinary() (data []byte, err error) {
	return json.Marshal(v)
}

func (v *Video) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, v)
}

var _ encoding.BinaryMarshaler = new(Video)
var _ encoding.BinaryUnmarshaler = new(Video)

func (v *Video) GetCacheKey() string {
	return "video_" + strconv.FormatInt(v.ID, 10)
}

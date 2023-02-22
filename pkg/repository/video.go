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
	ID        int64     `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Title     string `gorm:"type:varchar(50)"`
	AuthorId  int64  `gorm:"index"`
	PlayUrl   string `gorm:"type:varchar(100)"`
	CoverUrl  string `gorm:"type:varchar(100)"`
}

func (v *Video) TableName() string {
	return constants.VideoTableName
}

// MarshalBinary implements encoding.BinaryMarshaller for redis
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

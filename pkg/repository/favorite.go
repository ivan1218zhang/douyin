package repository

import (
	"douyin/pkg/constants"
	"encoding"
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type Favorite struct {
	Id        int64
	UserId    int64
	VideoId   int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (f *Favorite) TableName() string {
	return constants.FavoriteTableName
}

// MarshalBinary implements encoding.BinaryMarshaller for redis
func (f *Favorite) MarshalBinary() (data []byte, err error) {
	return json.Marshal(f)
}

func (f *Favorite) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, f)
}

var _ encoding.BinaryMarshaler = new(Favorite)
var _ encoding.BinaryUnmarshaler = new(Favorite)

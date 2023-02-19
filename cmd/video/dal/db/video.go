package db

import (
	"context"
	"douyin/kitex_gen/common"
	"douyin/pkg/repository"
	"gorm.io/gorm/clause"
	"time"
)

// MGetPublish query list of video info
func MGetPublish(ctx context.Context, userID int64) ([]*common.Video, error) {
	var res []*common.Video
	err := DB.WithContext(ctx).Table("video").Where("author_id = ?", userID).Find(&res).Error
	if err != nil {
		return res, err
	}
	return res, nil
}

// CreateVideo create video
func CreateVideo(ctx context.Context, video *repository.Video) error {
	return DB.WithContext(ctx).Table("video").Create(video).Error
}

// GetVideo get video by id
func GetVideo(ctx context.Context, id int64) (repository.Video, error) {
	var v repository.Video
	err := DB.WithContext(ctx).First(&v, id).Error
	return v, err
}

// MGetVideo multiple get list of Video info
func MGetVideo(ctx context.Context, latestTime int64) ([]*common.Video, error) {
	var res []*common.Video
	tm := time.Unix(0, latestTime*int64(time.Millisecond))
	if err := DB.WithContext(ctx).Table("video").Where("created_at < ?", tm.Format("2006-01-02 15:04:05")).Order(clause.OrderByColumn{Column: clause.Column{Name: "created_at"}, Desc: true}).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}
/*
// MGetVideo multiple get list of Video info
func MGetVideo(ctx context.Context, latestTime int64) (vs []*repository.Video, err error) {
	vs = []*repository.Video{}
	if latestTime != 0 {
		err = DB.WithContext(ctx).Where("created_at < ?", time.Unix(latestTime, 0)).Order("created_at DESC").Limit(10).Find(&vs).Error
	} else {
		err = DB.WithContext(ctx).Order("created_at DESC").Limit(10).Find(&vs).Error
	}

	return
}
 */
func GetVideoCreatedAt(ctx context.Context, videoId int64) (int64, error) {
	var res *repository.User
	err := DB.WithContext(ctx).Table("video").Where("id = ?", videoId).Select("created_at").First(res).Error
	if res == nil {
		return time.Now().UnixMilli(), err
	}
	return res.CreatedAt.UnixMilli(), err
}

// MGetVideoById multiple get list of Video info by ID
func MGetVideoById(ctx context.Context, idList []int64) ([]*common.Video, error) {
	var res []*common.Video
	if err := DB.WithContext(ctx).Table("video").Where("id in ? ", idList).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

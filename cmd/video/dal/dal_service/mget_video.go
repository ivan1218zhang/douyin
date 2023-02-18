package dbservice

import (
	"context"
	"douyin/cmd/video/dal/db"
	"douyin/cmd/video/dal/rdb"
	"douyin/pkg/repository"
)

func getVideoByID(id int64) repository.Video {
	v := rdb.GetVideoById(id)
	if (v != repository.Video{}) {
		return v
	}
	v, err := db.GetVideo(context.Background(), id)
	if err != nil {
		return repository.Video{}
	}
	go rdb.SetVideo(v)
	return v
}

func mGetVideoIDByTime(latestTime int64) ([]int64, error) {

	return nil, nil
}

func MGetVideo(latestTime int64) {

}

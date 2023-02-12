package pack

import (
	"douyin/kitex_gen/common"
	"douyin/pkg/repository"
)

// Video pack video info
func Video(v *repository.Video) *common.Video {
	if v == nil {
		return nil
	}

	return &common.Video{
		Id: v.ID,
		Author: &common.User{
			Id: int64(v.AuthorId),
		},
		PlayUrl:  v.PlayUrl,
		CoverUrl: v.CoverUrl,
	}
}

// Videos pack list of video info
func Videos(vs []*repository.Video) []*common.Video {
	videos := make([]*common.Video, len(vs))
	for i, v := range vs {
		if n := Video(v); n != nil {
			videos[i] = n
		}
	}
	return videos
}

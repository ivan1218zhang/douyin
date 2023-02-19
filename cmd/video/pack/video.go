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
		Id:       v.ID,
		AuthorId: v.AuthorId,
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

func UserIds(vs []*common.Video) []int64 {
	uIds := make([]int64, len(vs))
	if len(vs) == 0 {
		return uIds
	}
	for i, v := range vs {
		uIds[i] = v.AuthorId
	}

	return uIds
}

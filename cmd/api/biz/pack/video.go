package pack

import (
	common2 "douyin/cmd/api/biz/model/common"
	"douyin/kitex_gen/common"
)

func VideoList(videos []*common.Video) []*common2.Video {
	res := make([]*common2.Video, len(videos))
	for index, video := range videos {
		res[index] = &common2.Video{
			ID:            video.Id,
			Author:        User(video.Author),
			PlayURL:       video.PlayUrl,
			CoverURL:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    video.IsFavorite,
			Title:         video.Title,
			AuthorID:      video.AuthorId,
		}
	}
	return res
}

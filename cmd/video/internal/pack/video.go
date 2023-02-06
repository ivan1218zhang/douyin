package pack

import (
	"douyin/cmd/video/internal/dal/db"
	"douyin/kitex_gen/video/video"
)

// Video pack video info
func Video(v *db.Video) *video.Video {
	if v == nil {
		return nil
	}

	return &video.Video{
		Id: v.ID,
		Author: &video.User{
			Id:   int64(v.Author.ID),
			Name: v.Author.UserName,
		},
		PlayUrl:       v.PlayUrl,
		CoverUrl:      v.CoverUrl,
		FavoriteCount: v.FavoriteCount,
		CommentCount:  v.CommentCount,
	}
}

// Videos pack list of video info
func Videos(vs []*db.Video) []*video.Video {
	videos := make([]*video.Video, len(vs))
	for i, v := range vs {
		if n := Video(v); n != nil {
			videos[i] = n
		}
	}
	return videos
}

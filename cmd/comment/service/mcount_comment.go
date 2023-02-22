package service

import (
	"context"
	"douyin/cmd/comment/dal/db"
	"douyin/cmd/comment/dal/redis"
	"douyin/kitex_gen/comment"
)

type MCountCommentService struct {
	ctx context.Context
}

func NewMCountCommentService(ctx context.Context) *MCountCommentService {
	return &MCountCommentService{ctx: ctx}
}

func (s *MCountCommentService) MCountComment(req *comment.MCountCommentReq) ([]int64, error) {
	counts := make([]int64, len(req.VideoIdList))

	for i, id := range req.VideoIdList {
		count, err := redis.GetCommentCount(s.ctx, id)
		if err == nil {
			counts[i] = count
		}

		count, err = db.GetCommentCountByVideoID(s.ctx, id)
		if err == nil {
			redis.SetCommentCount(s.ctx, id, count)
			counts[i] = count
		}

	}
	return counts, nil
}

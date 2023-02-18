package service

import (
	"context"
	"douyin/cmd/comment/dal/db"
	"douyin/cmd/comment/dal/redis"
	"douyin/kitex_gen/comment"
)

type CountCommentService struct {
	ctx context.Context
}

func NewCountCommentService(ctx context.Context) *CountCommentService {
	return &CountCommentService{ctx: ctx}
}

func (s *CountCommentService) CountComment(req *comment.CountCommentReq) (int64, error) {
	count, err := redis.GetCommentCount(s.ctx, req.VideoId)
	if err == nil {
		return count, nil
	}
	count, err = db.GetCommentCountByVideoID(s.ctx, req.VideoId)
	if err == nil {
		redis.SetCommentCount(s.ctx, req.VideoId, count)
		return count, nil
	}
	return 0, nil
}

package service

import (
	"context"
	"douyin/cmd/comment/dal/db"
	"douyin/cmd/comment/dal/redis"
	"douyin/cmd/comment/pack"
	"douyin/cmd/comment/rpc"
	"douyin/kitex_gen/comment"
	"douyin/kitex_gen/common"
	"douyin/kitex_gen/user"
)

type MGetCommentService struct {
	ctx context.Context
}

func NewMGetCommentService(ctx context.Context) *MGetCommentService {
	return &MGetCommentService{ctx: ctx}
}
func (s *MGetCommentService) MGetComment(req *comment.MGetCommentReq) ([]*common.Comment, error) {
	// 查找缓存层是否存有
	comments, err := redis.GetCommentsForVideo(s.ctx, req.VideoId)
	if err == nil && comments != nil {
		return comments, nil
	}
	// Comments were not found in cache, fetch from database
	cs, err := db.GetCommentListByVideoId(s.ctx, req.VideoId)
	if err != nil {
		return nil, err
	}
	comments = pack.Comments(cs)
	//根据UserID获取用户信息
	idList := make([]int64, len(comments))
	for i, c := range comments {
		idList[i] = c.UserId
	}

	if len(idList) == 0 {
		// 存入缓存中
		if err := redis.SetComments(s.ctx, req.VideoId, comments); err != nil {
			return comments, err
		}
		return comments, nil
	}

	users, err := rpc.MGetUser(s.ctx, &user.MGetUserReq{IdList: idList, UserId: req.UserId})

	if err != nil {
		return nil, err
	}

	for i, c := range comments {
		c.User = users[i]
	}

	// 存入缓存中
	if err := redis.SetComments(s.ctx, req.VideoId, comments); err != nil {
		return comments, err
	}

	return comments, nil
}

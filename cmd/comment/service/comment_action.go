package service

import (
	"context"
	"douyin/cmd/comment/dal/db"
	"douyin/cmd/comment/pack"
	"douyin/cmd/comment/rpc"
	"douyin/kitex_gen/comment"
	"douyin/kitex_gen/common"
	"douyin/kitex_gen/user"
	"douyin/pkg/errno"
	"douyin/pkg/repository"
)

type CommentActionService struct {
	ctx context.Context
}

func NewCommentActionService(ctx context.Context) *CommentActionService {
	return &CommentActionService{ctx: ctx}
}
func (s *CommentActionService) CommentAction(req *comment.CommentActionReq) (*common.Comment, error) {
	switch req.ActionType {
	case 1:
		return publishComment(s.ctx, &repository.Comment{
			UserId:  req.UserId,
			VideoId: req.VideoId,
			Content: req.CommentText,
		})
	case 2:
		return deleteComment(s.ctx, req.CommentId)
	default:
		return nil, errno.ParamErr
	}
}

func deleteComment(ctx context.Context, commentId int64) (*common.Comment, error) {
	if err := db.DeleteComment(ctx, commentId); err != nil {
		return nil, err
	}
	return nil, nil
}

func publishComment(ctx context.Context, comment *repository.Comment) (*common.Comment, error) {
	comment1, err := db.CreateComment(ctx, comment)
	if err != nil {
		return nil, err
	}
	c := pack.Comment(comment1)
	// TODO:根据UserId获取用户信息
	u, err := rpc.GetUser(ctx, &user.GetUserReq{
		Id:     comment.UserId,
		UserId: comment.UserId,
	})
	c.User = u
	return c, nil
}

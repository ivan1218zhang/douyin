package main

import (
	"context"
	"douyin/cmd/comment/service"
	"douyin/cmd/video/pack"
	"douyin/kitex_gen/comment"
	"douyin/pkg/errno"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CommentAction implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentAction(ctx context.Context, req *comment.CommentActionReq) (resp *comment.CommentActionResp, err error) {
	resp = new(comment.CommentActionResp)

	c, err := service.NewCommentActionService(ctx).CommentAction(req)

	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	if c != nil {
		// TODO:通过c.UserID获得用户信息
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.SetComment(c)

	return resp, nil
}

// CountComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CountComment(ctx context.Context, req *comment.CountCommentReq) (resp *comment.CountCommentResp, err error) {
	// TODO: Your code here...
	return
}

// MCountComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) MCountComment(ctx context.Context, req *comment.MCountCommentReq) (resp *comment.MCountCommentResp, err error) {
	// TODO: Your code here...
	return
}

// MGetComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) MGetComment(ctx context.Context, req *comment.MGetCommentReq) (resp *comment.MGetCommentResp, err error) {
	// TODO: Your code here...
	return
}

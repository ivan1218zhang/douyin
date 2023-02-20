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

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.SetComment(c)

	return resp, nil
}

// CountComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CountComment(ctx context.Context, req *comment.CountCommentReq) (resp *comment.CountCommentResp, err error) {
	resp = new(comment.CountCommentResp)

	count, err := service.NewCountCommentService(ctx).CountComment(req)

	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.SetCommentCount(count)

	return resp, nil
}

// MCountComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) MCountComment(ctx context.Context, req *comment.MCountCommentReq) (resp *comment.MCountCommentResp, err error) {
	resp = new(comment.MCountCommentResp)

	counts, err := service.NewMCountCommentService(ctx).MCountComment(req)

	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.SetCommentCountList(counts)

	return resp, nil
}

// MGetComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) MGetComment(ctx context.Context, req *comment.MGetCommentReq) (resp *comment.MGetCommentResp, err error) {
	resp = new(comment.MGetCommentResp)

	comments, err := service.NewMGetCommentService(ctx).MGetComment(req)

	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.SetCommentList(comments)

	return resp, nil
}

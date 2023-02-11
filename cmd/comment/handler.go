package main

import (
	"context"
	"douyin/kitex_gen/comment"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CommentAction implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentAction(ctx context.Context, req *comment.CommentActionReq) (resp *comment.CommentActionResp, err error) {
	// TODO: Your code here...
	return
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

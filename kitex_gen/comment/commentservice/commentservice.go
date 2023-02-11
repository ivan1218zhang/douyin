// Code generated by Kitex v0.3.2. DO NOT EDIT.

package commentservice

import (
	"context"
	"douyin/kitex_gen/comment"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return commentServiceServiceInfo
}

var commentServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "CommentService"
	handlerType := (*comment.CommentService)(nil)
	methods := map[string]kitex.MethodInfo{
		"CommentAction": kitex.NewMethodInfo(commentActionHandler, newCommentServiceCommentActionArgs, newCommentServiceCommentActionResult, false),
		"CountComment":  kitex.NewMethodInfo(countCommentHandler, newCommentServiceCountCommentArgs, newCommentServiceCountCommentResult, false),
		"MCountComment": kitex.NewMethodInfo(mCountCommentHandler, newCommentServiceMCountCommentArgs, newCommentServiceMCountCommentResult, false),
		"MGetComment":   kitex.NewMethodInfo(mGetCommentHandler, newCommentServiceMGetCommentArgs, newCommentServiceMGetCommentResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "comment",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.3.2",
		Extra:           extra,
	}
	return svcInfo
}

func commentActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*comment.CommentServiceCommentActionArgs)
	realResult := result.(*comment.CommentServiceCommentActionResult)
	success, err := handler.(comment.CommentService).CommentAction(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommentServiceCommentActionArgs() interface{} {
	return comment.NewCommentServiceCommentActionArgs()
}

func newCommentServiceCommentActionResult() interface{} {
	return comment.NewCommentServiceCommentActionResult()
}

func countCommentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*comment.CommentServiceCountCommentArgs)
	realResult := result.(*comment.CommentServiceCountCommentResult)
	success, err := handler.(comment.CommentService).CountComment(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommentServiceCountCommentArgs() interface{} {
	return comment.NewCommentServiceCountCommentArgs()
}

func newCommentServiceCountCommentResult() interface{} {
	return comment.NewCommentServiceCountCommentResult()
}

func mCountCommentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*comment.CommentServiceMCountCommentArgs)
	realResult := result.(*comment.CommentServiceMCountCommentResult)
	success, err := handler.(comment.CommentService).MCountComment(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommentServiceMCountCommentArgs() interface{} {
	return comment.NewCommentServiceMCountCommentArgs()
}

func newCommentServiceMCountCommentResult() interface{} {
	return comment.NewCommentServiceMCountCommentResult()
}

func mGetCommentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*comment.CommentServiceMGetCommentArgs)
	realResult := result.(*comment.CommentServiceMGetCommentResult)
	success, err := handler.(comment.CommentService).MGetComment(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommentServiceMGetCommentArgs() interface{} {
	return comment.NewCommentServiceMGetCommentArgs()
}

func newCommentServiceMGetCommentResult() interface{} {
	return comment.NewCommentServiceMGetCommentResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CommentAction(ctx context.Context, req *comment.CommentActionReq) (r *comment.CommentActionResp, err error) {
	var _args comment.CommentServiceCommentActionArgs
	_args.Req = req
	var _result comment.CommentServiceCommentActionResult
	if err = p.c.Call(ctx, "CommentAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CountComment(ctx context.Context, req *comment.CountCommentReq) (r *comment.CountCommentResp, err error) {
	var _args comment.CommentServiceCountCommentArgs
	_args.Req = req
	var _result comment.CommentServiceCountCommentResult
	if err = p.c.Call(ctx, "CountComment", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MCountComment(ctx context.Context, req *comment.MCountCommentReq) (r *comment.MCountCommentResp, err error) {
	var _args comment.CommentServiceMCountCommentArgs
	_args.Req = req
	var _result comment.CommentServiceMCountCommentResult
	if err = p.c.Call(ctx, "MCountComment", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MGetComment(ctx context.Context, req *comment.MGetCommentReq) (r *comment.MGetCommentResp, err error) {
	var _args comment.CommentServiceMGetCommentArgs
	_args.Req = req
	var _result comment.CommentServiceMGetCommentResult
	if err = p.c.Call(ctx, "MGetComment", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

// Code generated by Kitex v0.3.2. DO NOT EDIT.

package relationservice

import (
	"context"
	"douyin/kitex_gen/relation"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	RelationAction(ctx context.Context, req *relation.RelationActionReq, callOptions ...callopt.Option) (r *relation.RelationActionResp, err error)
	MGetFollow(ctx context.Context, req *relation.MGetFollowReq, callOptions ...callopt.Option) (r *relation.MGetFollowResp, err error)
	MGetFollower(ctx context.Context, req *relation.MGetFollowerReq, callOptions ...callopt.Option) (r *relation.MGetFollowerResp, err error)
	CountFollow(ctx context.Context, req *relation.CountFollowReq, callOptions ...callopt.Option) (r *relation.CountFollowResp, err error)
	CountFollower(ctx context.Context, req *relation.CountFollowerReq, callOptions ...callopt.Option) (r *relation.CountFollowerResp, err error)
	IsFollow(ctx context.Context, req *relation.IsFollowReq, callOptions ...callopt.Option) (r *relation.IsFollowResp, err error)
	MCountFollow(ctx context.Context, req *relation.MCountFollowReq, callOptions ...callopt.Option) (r *relation.MCountFollowResp, err error)
	MCountFollower(ctx context.Context, req *relation.MCountFollowerReq, callOptions ...callopt.Option) (r *relation.MCountFollowerResp, err error)
	MIsFollow(ctx context.Context, req *relation.MIsFollowReq, callOptions ...callopt.Option) (r *relation.MIsFollowResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kRelationServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kRelationServiceClient struct {
	*kClient
}

func (p *kRelationServiceClient) RelationAction(ctx context.Context, req *relation.RelationActionReq, callOptions ...callopt.Option) (r *relation.RelationActionResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RelationAction(ctx, req)
}

func (p *kRelationServiceClient) MGetFollow(ctx context.Context, req *relation.MGetFollowReq, callOptions ...callopt.Option) (r *relation.MGetFollowResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MGetFollow(ctx, req)
}

func (p *kRelationServiceClient) MGetFollower(ctx context.Context, req *relation.MGetFollowerReq, callOptions ...callopt.Option) (r *relation.MGetFollowerResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MGetFollower(ctx, req)
}

func (p *kRelationServiceClient) CountFollow(ctx context.Context, req *relation.CountFollowReq, callOptions ...callopt.Option) (r *relation.CountFollowResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CountFollow(ctx, req)
}

func (p *kRelationServiceClient) CountFollower(ctx context.Context, req *relation.CountFollowerReq, callOptions ...callopt.Option) (r *relation.CountFollowerResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CountFollower(ctx, req)
}

func (p *kRelationServiceClient) IsFollow(ctx context.Context, req *relation.IsFollowReq, callOptions ...callopt.Option) (r *relation.IsFollowResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.IsFollow(ctx, req)
}

func (p *kRelationServiceClient) MCountFollow(ctx context.Context, req *relation.MCountFollowReq, callOptions ...callopt.Option) (r *relation.MCountFollowResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MCountFollow(ctx, req)
}

func (p *kRelationServiceClient) MCountFollower(ctx context.Context, req *relation.MCountFollowerReq, callOptions ...callopt.Option) (r *relation.MCountFollowerResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MCountFollower(ctx, req)
}

func (p *kRelationServiceClient) MIsFollow(ctx context.Context, req *relation.MIsFollowReq, callOptions ...callopt.Option) (r *relation.MIsFollowResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MIsFollow(ctx, req)
}
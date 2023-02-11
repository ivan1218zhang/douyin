// Code generated by Kitex v0.3.2. DO NOT EDIT.

package favoriteservice

import (
	"context"
	"douyin/kitex_gen/favorite"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	IsFavorite(ctx context.Context, req *favorite.IsFavoriteReq, callOptions ...callopt.Option) (r *favorite.IsFavoriteResp, err error)
	MIsFavorite(ctx context.Context, req *favorite.MIsFavoriteReq, callOptions ...callopt.Option) (r *favorite.MIsFavoriteResp, err error)
	CountFavorite(ctx context.Context, req *favorite.CountFavoriteReq, callOptions ...callopt.Option) (r *favorite.CountFavoriteResp, err error)
	MCountFavorite(ctx context.Context, req *favorite.MCountFavoriteReq, callOptions ...callopt.Option) (r *favorite.MCountFavoriteResp, err error)
	FavoriteAction(ctx context.Context, req *favorite.FavoriteActionReq, callOptions ...callopt.Option) (r *favorite.FavoriteActionResp, err error)
	MGetFavoriteVideo(ctx context.Context, req *favorite.MGetFavoriteVideoReq, callOptions ...callopt.Option) (r *favorite.MGetFavoriteVideoResp, err error)
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
	return &kFavoriteServiceClient{
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

type kFavoriteServiceClient struct {
	*kClient
}

func (p *kFavoriteServiceClient) IsFavorite(ctx context.Context, req *favorite.IsFavoriteReq, callOptions ...callopt.Option) (r *favorite.IsFavoriteResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.IsFavorite(ctx, req)
}

func (p *kFavoriteServiceClient) MIsFavorite(ctx context.Context, req *favorite.MIsFavoriteReq, callOptions ...callopt.Option) (r *favorite.MIsFavoriteResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MIsFavorite(ctx, req)
}

func (p *kFavoriteServiceClient) CountFavorite(ctx context.Context, req *favorite.CountFavoriteReq, callOptions ...callopt.Option) (r *favorite.CountFavoriteResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CountFavorite(ctx, req)
}

func (p *kFavoriteServiceClient) MCountFavorite(ctx context.Context, req *favorite.MCountFavoriteReq, callOptions ...callopt.Option) (r *favorite.MCountFavoriteResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MCountFavorite(ctx, req)
}

func (p *kFavoriteServiceClient) FavoriteAction(ctx context.Context, req *favorite.FavoriteActionReq, callOptions ...callopt.Option) (r *favorite.FavoriteActionResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FavoriteAction(ctx, req)
}

func (p *kFavoriteServiceClient) MGetFavoriteVideo(ctx context.Context, req *favorite.MGetFavoriteVideoReq, callOptions ...callopt.Option) (r *favorite.MGetFavoriteVideoResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MGetFavoriteVideo(ctx, req)
}
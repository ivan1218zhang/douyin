package rpc

import (
	"context"
	"douyin/cmd/api/biz/model/api"
	"douyin/cmd/api/biz/pack"
	"douyin/kitex_gen/message"
	"douyin/kitex_gen/message/messageservice"
	"douyin/pkg/constants"
	"github.com/cloudwego/kitex/client"
	"time"
)

var messageClient messageservice.Client

func initMessage() {
	c, err := messageservice.NewClient(
		"comment",
		client.WithHostPorts(constants.MessageServiceWithHostPorts),
		client.WithRPCTimeout(3*time.Second),
	)
	if err != nil {
		panic(err)
	}
	messageClient = c
}

// MessageAction 发消息
func MessageAction(ctx context.Context, req *message.MessageActionReq) (*api.MessageActionResp, error) {
	resp, err := messageClient.MessageAction(ctx, req)
	if err != nil {
		return nil, err
	}
	return &api.MessageActionResp{
		StatusCode:    resp.BaseResp.StatusCode,
		StatusMessage: resp.BaseResp.StatusMessage,
	}, nil
}

// MGetChatMessage 获取消息
func MGetChatMessage(ctx context.Context, req *message.MessageChatReq) (*api.MGetChatMessageResp, error) {
	resp, err := messageClient.GetMessageChat(ctx, req)

	if err != nil {

		return nil, err

	}

	return &api.MGetChatMessageResp{
		StatusCode:    resp.BaseResp.StatusCode,
		StatusMessage: resp.BaseResp.StatusMessage,
		MessageList:   pack.MessageList(resp.MessageList),
	}, nil
}

package main

import (
	"context"
	"douyin/cmd/message/pack"
	"douyin/cmd/message/service"
	message "douyin/kitex_gen/message"
	"douyin/pkg/errno"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// GetMessageChat implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) GetMessageChat(ctx context.Context, req *message.MessageChatReq) (resp *message.MessageChatResp, err error) {
	// TODO: Your code here...
	var MassageList []*message.Message
	resp = new(message.MessageChatResp)
	MassageList, err = service.NewGetMessageChatService(ctx).GetMessageChat(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		resp.MessageList = MassageList
		return resp, err
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)

	resp.MessageList = MassageList
	return resp, err
}

// MessageAction implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageAction(ctx context.Context, req *message.MessageActionReq) (resp *message.MessageActionResp, err error) {
	// TODO: Your code here...
	err = service.NewMessageActionService(ctx).MessageAction(req)
	resp = new(message.MessageActionResp)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, err
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, err
}

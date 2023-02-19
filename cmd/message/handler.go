package main

import (
	"context"
	"douyin/cmd/message/service"
	"douyin/cmd/user/pack"
	message "douyin/kitex_gen/message"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// GetmessageChat implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) GetmessageChat(ctx context.Context, req *message.MessageChatReq) (resp *message.MessageChatResp, err error) {
	// TODO: Your code here...
	var MassageList []*message.Message
	MassageList, err = service.NewGetMessageChatService(ctx).GetMessageChat(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		resp.MessageList = MassageList
		return resp, err
	}
	resp.BaseResp = pack.BuildBaseResp(err)
	resp.MessageList = MassageList
	return resp, err
}

// MessageAction implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageAction(ctx context.Context, req *message.MessageActionReq) (resp *message.MessageActionResp, err error) {
	// TODO: Your code here...
	err = service.NewMessageActionService(ctx).MessageAction(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, err
	}
	resp.BaseResp = pack.BuildBaseResp(err)
	return resp, err
}

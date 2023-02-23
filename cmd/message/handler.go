package main

import (
	"context"
	"douyin/cmd/message/pack"
	"douyin/cmd/message/service"
	"douyin/kitex_gen/message"
	"douyin/pkg/errno"
	"fmt"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// GetMassageChat implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) GetMassageChat(ctx context.Context, req *message.MessageChatReq) (resp *message.MessageChatResp, err error) {
	// TODO: Your code here...
	var MassageList []*message.Message
	resp = new(message.MessageChatResp)

	MassageList, err = service.NewGetMessageChatService(ctx).GetMessageChat(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		resp.MassageList = MassageList
		return resp, err
	}
	resp.BaseResp = pack.BuildBaseResp(err)
	resp.MassageList = MassageList
	fmt.Println(resp)
	return resp, err
}

// MassageAction implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MassageAction(ctx context.Context, req *message.MassageActionReq) (resp *message.MassageActionResp, err error) {
	// TODO: Your code here...
	err = service.NewMessageActionService(ctx).MessageAction(req)
	resp = new(message.MassageActionResp)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, err
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, err
}

// GetLatestMessage implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) GetLatestMessage(ctx context.Context, req *message.GetLatestMessageReq) (resp *message.GetLatestMessageResp, err error) {
	// TODO: Your code here...
	resp = new(message.GetLatestMessageResp)
	resp.Message = message.NewMessage()
	var cur message.Message
	cur, err = service.NewGetLatestMessageService(ctx).GetLatestMessage(req)

	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, err
	}

	resp.Message.FromUserId = cur.FromUserId
	resp.Message.ToUserId = cur.ToUserId
	resp.Message.Id = cur.Id
	resp.Message.Content = cur.Content
	resp.Message.CreateTime = cur.CreateTime
	resp.BaseResp = pack.BuildBaseResp(errno.Success)

	return resp, err

}

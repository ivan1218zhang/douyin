package service

import (
	"context"
	"douyin/cmd/message/dal/db"
	"douyin/kitex_gen/message"
)

type GetMessageChatService struct {
	ctx context.Context
}

// NewGetMessageChatService get message list service
func NewGetMessageChatService(ctx context.Context) *GetMessageChatService {
	return &GetMessageChatService{ctx: ctx}
}

// GetMessageChat get message list
func (s *GetMessageChatService) GetMessageChat(req *message.MessageChatReq) ([]*message.Message, error) {
	var res []*message.Message
	cur, err := db.GetMessageList(s.ctx, req.FromUserId, req.UserId)
	if err != nil {
		return res, err
	}
	for i := range cur {
		var temp message.Message
		temp.ToUserId = cur[i].ToUserID
		temp.Content = cur[i].Content
		temp.Id = cur[i].ID
		temp.CreateTime = cur[i].Date
		temp.FromUserId = cur[i].FromUserID
		res = append(res, &temp)
	}

	return res, nil

}

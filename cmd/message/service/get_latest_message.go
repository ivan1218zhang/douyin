package service

import (
	"context"
	"douyin/cmd/message/dal/db"
	"douyin/kitex_gen/message"
)

type GetLatestMessageService struct {
	ctx context.Context
}

// NewGetLatestMessageService get message list service
func NewGetLatestMessageService(ctx context.Context) *GetLatestMessageService {
	return &GetLatestMessageService{ctx: ctx}
}

// GetLatestMessage get message list
func (s *GetLatestMessageService) GetLatestMessage(req *message.GetLatestMessageReq) (message.Message, error) {
	var res message.Message
	res, err := db.GetLatestMessage(s.ctx, req.FromUserId, req.UserId)
	if err != nil {
		return res, err
	}
	return res, nil
}

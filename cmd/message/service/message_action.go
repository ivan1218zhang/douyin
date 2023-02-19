package service

import (
	"context"
	"douyin/cmd/message/dal/db"
	"douyin/kitex_gen/message"
)

type messageActionService struct {
	ctx context.Context
}

// NewMessageActionService new messageActionService
func NewMessageActionService(ctx context.Context) *messageActionService {
	return &messageActionService{ctx: ctx}
}

// MessageAction send message
func (s *messageActionService) MessageAction(req *message.MessageActionReq) error {

	err := db.SendMessage(s.ctx, req.UserId, req.ToUserId, req.Content)

	if err != nil {
		return err
	}
	return nil

}

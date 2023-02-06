package service

import (
	"context"
	"douyin/cmd/user/dal/db"
	"douyin/kitex_gen/user"
)

type GetUserByIdService struct {
	ctx context.Context
}

func NewGetUserByIdService(ctx context.Context) *GetUserByIdService {
	return &GetUserByIdService{ctx: ctx}
}

func (s *GetUserByIdService) GetUserById(req *user.GetUserByIdReq) (user *db.User, err error) {
	users, err := db.GetUsers(s.ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return users, nil
}

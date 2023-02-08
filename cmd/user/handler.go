package main

import (
	"context"
	"douyin/cmd/user/pack"
	"douyin/cmd/user/service"
	"douyin/kitex_gen/user"
	"douyin/pkg/errno"
	"douyin/pkg/repository"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *user.CreateUserReq) (resp *user.CreateUserResp, err error) {

	resp = new(user.CreateUserResp)

	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	user_id, err := service.NewCreateUserService(ctx).CreateUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.SetUserId(user_id)
	return resp, nil
}

// GetUserById implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserById(ctx context.Context, req *user.GetUserByIdReq) (resp *user.GetUserByIdResp, err error) {

	resp = new(user.GetUserByIdResp)
	users := &repository.User{}

	users, err = service.NewGetUserByIdService(ctx).GetUserById(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	if users == nil {
		resp.BaseResp = pack.BuildBaseResp(errno.Success)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *user.CheckUserReq) (resp *user.CheckUserResp, err error) {
	// TODO: Your code here...
	return
}

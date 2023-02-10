package service_test

import (
	"context"
	"douyin/cmd/user/dal"
	"douyin/cmd/user/service"
	"douyin/kitex_gen/user"
	"testing"
)

func TestGetUser(t *testing.T) {

	dal.Init()
	ctx := context.Background()
	create_service := service.NewCreateUserService(ctx)

	// Test case 1: check if user is created successfully
	create_req := &user.CreateUserReq{
		Username: "testuser_gettest",
		Password: "password",
	}
	var user_id int64
	var err error
	if user_id, err = create_service.CreateUser(create_req); err != nil {
		t.Errorf("CreateUser returned an error: %v", err)
	}

	ctx = context.Background()
	service := service.NewGetUserService(ctx)

	req1 := &user.GetUserReq{
		Id: user_id,
	}

	if u, err := service.GetUser(req1); err != nil && u.Name == "testuser_gettest" {
		t.Errorf("GetUser returned an error: %v", err)
	}
}

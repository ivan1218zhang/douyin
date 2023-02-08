package service_test

import (
	"context"
	"douyin/cmd/user/dal"
	"douyin/cmd/user/service"
	"douyin/kitex_gen/user"
	"testing"
)

func TestCheckUser(t *testing.T) {

	dal.Init()
	ctx := context.Background()
	create_service := service.NewCreateUserService(ctx)

	// Test case 1: check if user is created successfully
	create_req := &user.CreateUserReq{
		Username: "testuser",
		Password: "password",
	}
	if _, err := create_service.CreateUser(create_req); err != nil {
		t.Errorf("CreateUser returned an error: %v", err)
	}

	ctx = context.Background()
	check_service := service.NewCheckUserService(ctx)

	check_req1 := &user.CheckUserReq{
		Username: "testuser",
		Password: "password",
	}
	if _, err := check_service.CheckUser(check_req1); err != nil {
		t.Errorf("CheckUser returned an error: %v", err)
	}

	check_req2 := &user.CheckUserReq{
		Username: "testuser",
		Password: "password1",
	}
	if _, err := check_service.CheckUser(check_req2); err == nil {
		t.Errorf("CheckUser is wrong")
	}
}

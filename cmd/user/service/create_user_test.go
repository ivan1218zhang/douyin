package service_test

import (
	"context"
	"douyin/cmd/user/dal"
	"douyin/cmd/user/service"
	"douyin/kitex_gen/user"
	"douyin/pkg/errno"
	"testing"
)

func TestCreateUser(t *testing.T) {
	dal.Init()
	ctx := context.Background()
	service := service.NewCreateUserService(ctx)

	// Test case 1: check if user is created successfully
	req := &user.CreateUserRequest{
		UserName: "testuser1",
		Password: "password",
	}
	if err := service.CreateUser(req); err != nil {
		t.Errorf("CreateUser returned an error: %v", err)
	}

	// Test case 2: check if user already exists error is returned
	req = &user.CreateUserRequest{
		UserName: "testuser1",
		Password: "password",
	}
	if err := service.CreateUser(req); err != errno.UserAlreadyExistErr {
		t.Errorf("CreateUser should have returned UserAlreadyExistErr error")
	}
}

package service_test

import (
	"context"
	"douyin/cmd/user/dal"
	"douyin/cmd/user/service"
	"douyin/kitex_gen/user"
	"fmt"
	"testing"
)

func TestMGetUser(t *testing.T) {

	dal.Init()
	ctx := context.Background()
	create_service := service.NewCreateUserService(ctx)

	// Test case 1: check if user is created successfully
	create_req1 := &user.CreateUserReq{
		Username: "mgetuser0",
		Password: "password",
	}
	create_req2 := &user.CreateUserReq{
		Username: "mgetuser1",
		Password: "password",
	}
	var user_id1, user_id2 int64
	var err error

	if user_id1, err = create_service.CreateUser(create_req1); err != nil {
		t.Errorf("CreateUser returned an error: %v", err)
	}
	if user_id2, err = create_service.CreateUser(create_req2); err != nil {
		t.Errorf("CreateUser returned an error: %v", err)
	}

	ctx = context.Background()
	service := service.NewMGetUserService(ctx)

	req := &user.MGetUserReq{
		IdList: []int64{user_id1, user_id2},
	}

	user_list, err := service.MGetUser(req)

	if err != nil {
		t.Errorf("GetUser returned an error: %v", err)
	}

	for i, u := range user_list {
		if u.Name != fmt.Sprint("mgetuser", i) {
			t.Error("GetUser error", u.Name, fmt.Sprint("mgetuser", i))
		}
	}
}

package main

import (
	"context"
	"douyin/cmd/user/dal"
	"douyin/cmd/user/pack"
	"douyin/kitex_gen/user"
	"douyin/pkg/errno"
	"testing"
)

func TestCreateUser(t *testing.T) {
	tests := []struct {
		req      *user.CreateUserRequest
		expected *user.CreateUserResponse
	}{
		{
			req: &user.CreateUserRequest{
				UserName: "user1",
				Password: "password1",
			},
			expected: &user.CreateUserResponse{
				BaseResp: pack.BuildBaseResp(errno.Success),
			},
		},
		{
			req: &user.CreateUserRequest{
				UserName: "",
				Password: "password2",
			},
			expected: &user.CreateUserResponse{
				BaseResp: pack.BuildBaseResp(errno.ParamErr),
			},
		},
		{
			req: &user.CreateUserRequest{
				UserName: "user3",
				Password: "",
			},
			expected: &user.CreateUserResponse{
				BaseResp: pack.BuildBaseResp(errno.ParamErr),
			},
		},
	}

	dal.Init()

	for _, test := range tests {
		resp, err := new(UserServiceImpl).CreateUser(context.Background(), test.req)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.BaseResp.StatusCode != test.expected.BaseResp.StatusCode {
			t.Fatalf("expected %d but got %d", test.expected.BaseResp.StatusCode, resp.BaseResp.StatusCode)
		}
	}
}

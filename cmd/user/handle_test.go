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
		req      *user.CreateUserReq
		expected *user.CreateUserResp
	}{
		{
			req: &user.CreateUserReq{
				Username: "user1",
				Password: "password1",
			},
			expected: &user.CreateUserResp{
				BaseResp: pack.BuildBaseResp(errno.Success),
			},
		},
		{
			req: &user.CreateUserReq{
				Username: "",
				Password: "password2",
			},
			expected: &user.CreateUserResp{
				BaseResp: pack.BuildBaseResp(errno.ParamErr),
			},
		},
		{
			req: &user.CreateUserReq{
				Username: "user3",
				Password: "",
			},
			expected: &user.CreateUserResp{
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

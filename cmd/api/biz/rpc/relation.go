package rpc

import (
	"context"
	"douyin/cmd/api/biz/model/api"
	"douyin/cmd/api/biz/pack"
	"douyin/kitex_gen/message"
	"douyin/kitex_gen/relation"
	"douyin/kitex_gen/relation/relationservice"
	"douyin/pkg/constants"
	"github.com/cloudwego/kitex/client"
	"time"
)

var relationClient relationservice.Client

func initRelation() {
	c, err := relationservice.NewClient(
		"relation",
		client.WithHostPorts(constants.RelationServiceWithHostPorts),
		client.WithRPCTimeout(3*time.Second),
	)
	if err != nil {
		panic(err)
	}
	relationClient = c
}

// RelationAction 关注操作
func RelationAction(ctx context.Context, req *relation.RelationActionReq) (*api.RelationActionResp, error) {
	resp, err := relationClient.RelationAction(ctx, req)
	if err != nil {
		return nil, err
	}
	return &api.RelationActionResp{
		StatusCode:    resp.BaseResp.StatusCode,
		StatusMessage: resp.BaseResp.StatusMessage,
	}, nil
}

// MGetFollow 关注列表
func MGetFollow(ctx context.Context, req *relation.MGetFollowReq) (*api.MGetFollowResp, error) {
	resp, err := relationClient.MGetFollow(ctx, req)
	if err != nil {
		return nil, err
	}
	return &api.MGetFollowResp{
		StatusCode:    resp.BaseResp.StatusCode,
		StatusMessage: resp.BaseResp.StatusMessage,
		UserList:      pack.UserList(resp.UserList),
	}, nil
}

// MGetFollower 粉丝列表
func MGetFollower(ctx context.Context, req *relation.MGetFollowerReq) (*api.MGetFollowerResp, error) {
	resp, err := relationClient.MGetFollower(ctx, req)
	if err != nil {
		return nil, err
	}
	return &api.MGetFollowerResp{
		StatusCode:    resp.BaseResp.StatusCode,
		StatusMessage: resp.BaseResp.StatusMessage,
		UserList:      pack.UserList(resp.UserList),
	}, nil
}

type FriendUser struct {
	Id            int64  `thrift:"id,1" json:"id"`
	Name          string `thrift:"name,2" json:"name"`
	FollowCount   int64  `thrift:"follow_count,3" json:"follow_count"`
	FollowerCount int64  `thrift:"follower_count,4" json:"follower_count"`
	IsFollow      bool   `thrift:"is_follow,5" json:"is_follow"`
	Message       string `json:"message"`
	MSGType       int64  `json:"msgType"`
}
type MGetFriendResp1 struct {
	StatusCode     int64        `thrift:"status_code,1" json:"status_code"`
	StatusMessage  string       `thrift:"status_message,2" json:"status_message"`
	FriendUserList []FriendUser `json:"user_list"`
}

// MGetFriend 朋友列表
func MGetFriend(ctx context.Context, req *relation.MGetFriendReq) (*MGetFriendResp1, error) {
	resp, err := relationClient.MGetFriend(ctx, req)
	if err != nil {
		return nil, err
	}
	FriendUserList := make([]FriendUser, len(resp.UserList))
	for i := range resp.UserList {
		FriendUserList[i].Id = resp.UserList[i].Id
		FriendUserList[i].FollowCount = resp.UserList[i].FollowCount
		FriendUserList[i].Name = resp.UserList[i].Name
		FriendUserList[i].IsFollow = resp.UserList[i].IsFollow
		FriendUserList[i].FollowerCount = resp.UserList[i].FollowerCount
		req1 := new(message.GetLatestMessageReq)
		req1.UserId = req.UserId
		req1.FromUserId = resp.UserList[i].Id

		resp1, err := messageClient.GetLatestMessage(ctx, req1)
		if err != nil {
			return nil, err
		}
		if resp1 != nil {
			FriendUserList[i].Message = resp1.Message.Content
		} else {
			FriendUserList[i].Message = "你们还没有发过消息呢"
		}
		if resp1.Message.FromUserId == req.UserId {
			FriendUserList[i].MSGType = 1
		} else {
			FriendUserList[i].MSGType = 0
		}
	}

	return &MGetFriendResp1{
		StatusCode:     resp.BaseResp.StatusCode,
		StatusMessage:  resp.BaseResp.StatusMessage,
		FriendUserList: FriendUserList,
	}, nil
}

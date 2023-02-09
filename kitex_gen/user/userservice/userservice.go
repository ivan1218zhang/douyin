// Code generated by Kitex v0.3.2. DO NOT EDIT.

package userservice

import (
	"context"
	"douyin/kitex_gen/user"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

var userServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*user.UserService)(nil)
	methods := map[string]kitex.MethodInfo{
		"CreateUser": kitex.NewMethodInfo(createUserHandler, newUserServiceCreateUserArgs, newUserServiceCreateUserResult, false),
		"GetUser":    kitex.NewMethodInfo(getUserHandler, newUserServiceGetUserArgs, newUserServiceGetUserResult, false),
		"CheckUser":  kitex.NewMethodInfo(checkUserHandler, newUserServiceCheckUserArgs, newUserServiceCheckUserResult, false),
		"MGetUser":   kitex.NewMethodInfo(mGetUserHandler, newUserServiceMGetUserArgs, newUserServiceMGetUserResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "user",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.3.2",
		Extra:           extra,
	}
	return svcInfo
}

func createUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceCreateUserArgs)
	realResult := result.(*user.UserServiceCreateUserResult)
	success, err := handler.(user.UserService).CreateUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceCreateUserArgs() interface{} {
	return user.NewUserServiceCreateUserArgs()
}

func newUserServiceCreateUserResult() interface{} {
	return user.NewUserServiceCreateUserResult()
}

func getUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceGetUserArgs)
	realResult := result.(*user.UserServiceGetUserResult)
	success, err := handler.(user.UserService).GetUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceGetUserArgs() interface{} {
	return user.NewUserServiceGetUserArgs()
}

func newUserServiceGetUserResult() interface{} {
	return user.NewUserServiceGetUserResult()
}

func checkUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceCheckUserArgs)
	realResult := result.(*user.UserServiceCheckUserResult)
	success, err := handler.(user.UserService).CheckUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceCheckUserArgs() interface{} {
	return user.NewUserServiceCheckUserArgs()
}

func newUserServiceCheckUserResult() interface{} {
	return user.NewUserServiceCheckUserResult()
}

func mGetUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceMGetUserArgs)
	realResult := result.(*user.UserServiceMGetUserResult)
	success, err := handler.(user.UserService).MGetUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceMGetUserArgs() interface{} {
	return user.NewUserServiceMGetUserArgs()
}

func newUserServiceMGetUserResult() interface{} {
	return user.NewUserServiceMGetUserResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateUser(ctx context.Context, req *user.CreateUserReq) (r *user.CreateUserResp, err error) {
	var _args user.UserServiceCreateUserArgs
	_args.Req = req
	var _result user.UserServiceCreateUserResult
	if err = p.c.Call(ctx, "CreateUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUser(ctx context.Context, req *user.GetUserReq) (r *user.GetUserResp, err error) {
	var _args user.UserServiceGetUserArgs
	_args.Req = req
	var _result user.UserServiceGetUserResult
	if err = p.c.Call(ctx, "GetUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CheckUser(ctx context.Context, req *user.CheckUserReq) (r *user.CheckUserResp, err error) {
	var _args user.UserServiceCheckUserArgs
	_args.Req = req
	var _result user.UserServiceCheckUserResult
	if err = p.c.Call(ctx, "CheckUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MGetUser(ctx context.Context, req *user.MGetUserReq) (r *user.MGetUserResp, err error) {
	var _args user.UserServiceMGetUserArgs
	_args.Req = req
	var _result user.UserServiceMGetUserResult
	if err = p.c.Call(ctx, "MGetUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

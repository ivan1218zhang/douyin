// Code generated by hertz generator.

package api

import (
	"context"
	"douyin/cmd/api/biz/rpc"
	"douyin/kitex_gen/user"
	"douyin/kitex_gen/video"
	"io/ioutil"

	api "douyin/cmd/api/biz/model/api"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// MGetVideo .
// @router /douyin/feed/ [GET]
func MGetVideo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.MGetVideoReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp, err := rpc.MGetVideo(ctx, &video.MGetVideoReq{
		LatestTime: req.LatestTime,
		UserId:     0,
	})

	c.JSON(consts.StatusOK, resp)
}

// Register .
// @router /douyin/user/register/ [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.RegisterReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp, err := rpc.CreateUser(ctx, &user.CreateUserReq{
		Username: req.Username,
		Password: req.Password,
	})

	c.JSON(consts.StatusOK, resp)
}

// Login .
// @router /douyin/user/login/ [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.LoginReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.LoginResp)

	c.JSON(consts.StatusOK, resp)
}

// GetUser .
// @router /douyin/user/ [GET]
func GetUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.GetUserReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp, err := rpc.GetUser(ctx, &user.GetUserReq{Id: req.UserID})

	c.JSON(consts.StatusOK, resp)
}

// Publish .
// @router /douyin/publish/action/ [POST]
func Publish(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.PublishReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	file, err := c.FormFile("data")
	source, err := file.Open()
	defer source.Close()
	data, err := ioutil.ReadAll(source)
	resp, err := rpc.Publish(ctx, &video.PublishReq{
		UserId: 0,
		Title:  req.Title,
		Data:   data,
	})

	c.JSON(consts.StatusOK, resp)
}

// MGetPublish .
// @router /douyin/publish/list/ [GET]
func MGetPublish(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.MGetPublishReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	resp, err := rpc.MGetPublish(ctx, &video.MGetPublishReq{UserId: req.UserID})

	c.JSON(consts.StatusOK, resp)
}
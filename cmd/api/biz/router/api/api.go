// Code generated by hertz generator. DO NOT EDIT.

package Api

import (
	api "douyin/cmd/api/biz/handler/api"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_douyin := root.Group("/douyin", _douyinMw()...)
		{
			_feed := _douyin.Group("/feed", _feedMw()...)
			_feed.GET("/", append(_mgetvideoMw(), api.MGetVideo)...)
		}
		{
			_publish := _douyin.Group("/publish", _publishMw()...)
			{
				_action := _publish.Group("/action", _actionMw()...)
				_action.POST("/", append(_publish0Mw(), api.Publish)...)
			}
			{
				_list := _publish.Group("/list", _listMw()...)
				_list.GET("/", append(_mgetpublishMw(), api.MGetPublish)...)
			}
		}
		{
			_user := _douyin.Group("/user", _userMw()...)
			_user.GET("/", append(_getuserMw(), api.GetUser)...)
			{
				_login := _user.Group("/login", _loginMw()...)
				_login.POST("/", append(_login0Mw(), api.Login)...)
			}
			{
				_register := _user.Group("/register", _registerMw()...)
				_register.POST("/", append(_register0Mw(), api.Register)...)
			}
		}
	}
}
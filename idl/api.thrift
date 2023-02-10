namespace go api

include 'common.thrift'

struct MGetVideoReq {
    1:i64 latest_time
    2:string token
}

struct MGetVideoResp {
    1:i64 status_code
    2:string status_message
    3:list<common.Video> video_list
    4:i64 next_time
}

struct RegisterReq {
    1:string username
    2:string password
}

struct RegisterResp {
    1:i64 status_code
    2:string status_message
    3:i64 user_id
    4:string token
}

struct LoginReq {
    1:string username
    2:string password
}

struct LoginResp {
     1:i64 status_code
     2:string status_message
     3:i64 user_id
     4:string token
}

struct GetUserReq {
    1:i64 user_id
    2:string token
}

struct GetUserResp {
    1:i64 status_code
    2:string status_message
    3:common.User user
}

struct PublishReq {
    2:string token
    3:string title
}

struct PublishResp {
    1:i64 status_code
    2:string status_message
}

struct MGetPublishReq {
    1:string token
    2:i64 user_id
}

struct MGetPublishResp{
    1:i64 status_code
    2:string status_message
    3:list<common.Video> video_list
}

service ApiService {
    MGetVideoResp MGetVideo(1:MGetVideoReq req) (api.get="/douyin/feed/")
    RegisterResp Register(1:RegisterReq req) (api.post="/douyin/user/register/")
    LoginResp Login(1:LoginReq req) (api.post="/douyin/user/login/")
    GetUserResp GetUser(1:GetUserReq req) (api.get="/douyin/user/")
    PublishResp Publish(1:PublishReq req) (api.post="/douyin/publish/action/")
    MGetPublishResp MGetPublish(1:MGetPublishReq req) (api.get="/douyin/publish/list/")
}
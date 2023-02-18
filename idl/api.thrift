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

struct MGetPublishResp {
    1:i64 status_code
    2:string status_message
    3:list<common.Video> video_list
}

struct FavoriteActionReq {
    1:string token
    2:i64 video_id
    3:i32 action_type
}

struct FavoriteActionResp {
    1:i64 status_code
    2:string status_message
}

struct MGetFavoriteReq {
    1:i64 user_id
    2:string token
}

struct MGetFavoriteResp {
    1:i64 status_code
    2:string status_message
    3:list<common.Video> video_list
}

struct CommentActionReq {
    1:string token
    2:i64 video_id
    3:i32 action_type
    4:string comment_context
    5:i64 comment_id
}

struct CommentActionResp {
    1:i64 status_code
    2:string status_message
    3:common.Comment comment
}

struct MGetCommentReq {
    1:string token
    2:i64 video_id
}

struct MGetCommentResp {
    1:i64 status_code
    2:string status_message
    3:list<common.Comment> comment_list
}

struct RelationActionReq {
    1:string token
    2:i64 to_user_id
    3:i32 action_type
}

struct RelationActionResp {
    1:i64 status_code
    2:string status_message
}

struct MGetFollowReq {
    1:i64 user_id
    2:string token
}

struct MGetFollowResp {
    1:i64 status_code
    2:string status_message
    3:list<common.User> user_list
}

struct MGetFollowerReq {
    1:i64 user_id
    2:string token
}

struct MGetFollowerResp {
    1:i64 status_code
    2:string status_message
    3:list<common.User> user_list
}

service ApiService {
    MGetVideoResp MGetVideo(1:MGetVideoReq req) (api.get="/douyin/feed/")
    RegisterResp Register(1:RegisterReq req) (api.post="/douyin/user/register/")
    LoginResp Login(1:LoginReq req) (api.post="/douyin/user/login/")
    GetUserResp GetUser(1:GetUserReq req) (api.get="/douyin/user/")
    PublishResp Publish(1:PublishReq req) (api.post="/douyin/publish/action/")
    MGetPublishResp MGetPublish(1:MGetPublishReq req) (api.get="/douyin/publish/list/")
    FavoriteActionResp FavoriteAction(1:FavoriteActionReq req) (api.post="/douyin/favorite/action/")
    MGetFavoriteResp MGetFavorite(1:MGetFavoriteReq req) (api.get="/douyin/favorite/list/")
    CommentActionResp CommentAction(1:CommentActionReq req) (api.post="/douyin/comment/action/")
    MGetCommentResp MGetComment(1:MGetCommentReq req) (api.get="/douyin/comment/list/")
    RelationActionResp RelationAction(1:RelationActionReq req) (api.post="/douyin/relation/action/")
    MGetFollowResp MGetFollow(1:MGetFollowReq req) (api.get="/douyin/relation/follow/list/")
    MGetFollowerResp MGetFollower(1:MGetFollowerReq req) (api.get="/douyin/relation/follower/list/")
}
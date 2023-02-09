namespace go favorite

include 'common.thrift'

struct IsFavoriteReq{
    i64 user_id
    i64 video_id
}
struct IsFavoriteResp{
    1:common.BaseResp base_resp
    2:bool is_favorite
}

struct IsFavoriteListReq{
    1:i64 user_id
    2:list<i64> video_id
}

struct IsFavoriteListResp{
    1:common.BaseResp base_resp
    2:list<bool> is_favorite
}

struct FavoriteActionReq{
    1:i64 user_id
    2:i64 video_id
    3:i32 action_type
}

struct FavoriteActionResp{
    1:common.BaseResp base_resp
}

struct GetFavoritesByUserIdReq{
    1:i64 user_id
}

struct GetFavoritesByUserIdResp{
    1:common.BaseResp base_resp
    2:list<common.Video> favorites
}

struct CountFavoriteReq{
    1:i64 video_id
}

struct CountFavoriteResp{
    1:common.BaseResp base_resp
    2:i64 favorite_count
}

struct CountFavoriteListReq{
    1:list<i64> video_id
}

struct CountFavoriteListResp{
    1:common.BaseResp base_resp
    2:list<i64> favorite_count
}

service FavoriteService {
    IsFavoriteResp IsFavorite(1:IsFavoriteReq req)
    IsFavoriteListResp IsFavoriteList(1:IsFavoriteListReq req)
    CountFavoriteResp CountFavorite(1:CountFavoriteReq req)
    CountFavoriteListResp CountFavoriteList(1:CountFavoriteListReq req)
    FavoriteActionResp FavoriteAction(1:FavoriteActionReq req)
    GetFavoritesByUserIdResp GetFavoritesByUserId(1:GetFavoritesByUserIdReq req)
}

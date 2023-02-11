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

struct MIsFavoriteReq{
    1:i64 user_id
    2:list<i64> video_id_list
}

struct MIsFavoriteResp{
    1:common.BaseResp base_resp
    2:list<bool> is_favorite_list
}

struct FavoriteActionReq{
    1:i64 user_id
    2:i64 video_id
    3:i32 action_type
}

struct FavoriteActionResp{
    1:common.BaseResp base_resp
}

struct MGetFavoriteVideoReq{
    1:i64 user_id
}

struct MGetFavoriteVideoResp{
    1:common.BaseResp base_resp
    2:list<common.Video> video_list
}

struct CountFavoriteReq{
    1:i64 video_id
}

struct CountFavoriteResp{
    1:common.BaseResp base_resp
    2:i64 favorite_count
}

struct MCountFavoriteReq{
    1:list<i64> video_id_list
}

struct MCountFavoriteResp{
    1:common.BaseResp base_resp
    2:list<i64> favorite_count_list
}

service FavoriteService {
    IsFavoriteResp IsFavorite(1:IsFavoriteReq req)
    MIsFavoriteResp MIsFavorite(1:MIsFavoriteReq req)
    CountFavoriteResp CountFavorite(1:CountFavoriteReq req)
    MCountFavoriteResp MCountFavorite(1:MCountFavoriteReq req)
    FavoriteActionResp FavoriteAction(1:FavoriteActionReq req)
    MGetFavoriteVideoResp MGetFavoriteVideo(1:MGetFavoriteVideoReq req)
}

namespace go favorite

include 'common.thrift'

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

struct MCountFavoriteReq{
    1:list<i64> video_id_list
}

struct MCountFavoriteResp{
    1:common.BaseResp base_resp
    2:list<i64> favorite_count_list
}

service FavoriteService {
    MIsFavoriteResp MIsFavorite(1:MIsFavoriteReq req)
    MCountFavoriteResp MCountFavorite(1:MCountFavoriteReq req)
    FavoriteActionResp FavoriteAction(1:FavoriteActionReq req)
    MGetFavoriteVideoResp MGetFavoriteVideo(1:MGetFavoriteVideoReq req)
}

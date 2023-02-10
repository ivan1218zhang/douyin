namespace go relation

include 'common.thrift'

struct RelationActionReq{
    1:i64 user_id
    2:i64 to_user_id
    3:i32 action_type
}
struct RelationActionResp{
    1:common.BaseResp base_resp
}

struct MGetFollowReq{
    1:i64 user_id
}
struct MGetFollowResp{
    1:common.BaseResp base_resp
    2:list<common.User> user_list
}

struct MGetFollowerReq{
    1:i64 user_id
}
struct MGetFollowerResp{
    1: common.BaseResp base_resp
    2:list<i64> user_list
}

struct CountFollowReq{
    1:i64 user_id
}

struct CountFollowResp{
    1:common.BaseResp base_resp
    2:i64 follow_count
}

struct CountFollowerReq{
    1:i64 user_id
}

struct CountFollowerResp{
    1:common.BaseResp base_resp
    2:i64 follower_count
}

struct IsFollowReq{
    1:i64 user_id
    2:i64 to_user_id
}

struct IsFollowResp{
    1:common.BaseResp base_resp
    2:bool is_follow
}

struct MCountFollowReq{
    1:list<i64> user_id_list
}

struct MCountFollowResp{
    1:common.BaseResp base_resp
    2:list<i64> follow_count_list
}

struct MCountFollowerReq{
    1:list<i64> user_id_list
}

struct MCountFollowerResp{
    1:common.BaseResp base_resp
    2:list<i64> follower_count_list
}

struct MIsFollowReq{
    1:list<i64> user_id_list
    2:list<i64> to_user_id_list
}

struct MIsFollowResp{
    1:common.BaseResp base_resp
    2:list<bool> is_follow_list
}

service RelationService {
    RelationActionResp RelationAction(1:RelationActionReq req)
    MGetFollowResp MGetFollow(1:MGetFollowReq req)
    MGetFollowerResp MGetFollower(1:MGetFollowerReq req)
    CountFollowResp CountFollow(1:CountFollowReq req)
    CountFollowerResp CountFollower(1:CountFollowerReq req)
    IsFollowResp IsFollow(1:IsFollowReq req)
    MCountFollowResp MCountFollow(1:MCountFollowReq req)
    MCountFollowerResp MCountFollower(1:MCountFollowerReq req)
    MIsFollowResp MIsFollow(1:MIsFollowReq req)
}
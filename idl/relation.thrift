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
    2:i64 id
}
struct MGetFollowResp{
    1:common.BaseResp base_resp
    2:list<common.User> user_list
}

struct MGetFollowerReq{
    1:i64 user_id
    2:i64 id
}

struct MGetFollowerResp{
    1: common.BaseResp base_resp
    2:list<common.User> user_list
}

struct MGetFriendReq{
    1:i64 user_id
}

struct MGetFriendResp{
    1:common.BaseResp base_resp
    2:list<common.User> user_list
}

struct GetRelationReq{
    1:i64 user_id
    2:i64 to_user_id
}

struct GetRelationResp{
    1:common.BaseResp base_resp
    2:i32 follow_count
    3:i32 follower_count
    4:bool is_follow
}

struct MGetRelationReq{
    1:i64 user_id
    2:list<i64> to_user_id_list
}

struct MGetRelationResp{
    1:common.BaseResp base_resp
    2:list<i32> follow_count_list
    3:list<i32> follower_count_list
    4:list<bool> is_follow_list
}

service RelationService {
    RelationActionResp RelationAction(1:RelationActionReq req)
    MGetFollowResp MGetFollow(1:MGetFollowReq req)
    MGetFollowerResp MGetFollower(1:MGetFollowerReq req)
    MGetFriendResp MGetFriend(1:MGetFriendReq req)
    GetRelationResp GetRelation(1:GetRelationReq req)
    MGetRelationResp MGetRelation(1:MGetRelationReq req)
}
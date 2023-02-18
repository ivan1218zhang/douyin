include 'common.thrift'

namespace go user

struct CreateUserReq{
    1:string username
    2:string password
}

struct CreateUserResp{
    1:common.BaseResp base_resp
    2:i64 user_id
}

struct GetUserReq{
    1:i64 id
    2:i64 user_id
}

struct GetUserResp{
    1:common.BaseResp base_resp
    2:common.User user
}

struct CheckUserReq {
    1: string username
    2: string password
}

struct CheckUserResp {
    1: i64 user_id
    2: common.BaseResp base_resp
}

struct MGetUserReq{
    1:list<i64> id_list
    2:i64 user_id
}

struct MGetUserResp{
    1:common.BaseResp base_resp
    2:list<common.User> user_list
}

service UserService {
    CreateUserResp CreateUser(1:CreateUserReq req)
    GetUserResp GetUser(1:GetUserReq req)
    CheckUserResp CheckUser(1: CheckUserReq req)
    MGetUserResp MGetUser(1:MGetUserReq req)
}
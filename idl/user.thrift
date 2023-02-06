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

struct GetUserByIdReq{
    1:i64 id
}

struct GetUserByIdResp{
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

service UserService {
    CreateUserResp CreateUser(1:CreateUserReq req)
    GetUserByIdResp GetUserById(1:GetUserByIdReq req)
    CheckUserResp CheckUser(1: CheckUserReq req)
}
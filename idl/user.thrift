namespace go user

struct User{
    1:i64 id
    2:string name
    3:i64 follow_count
    4:i64 follower_count
    5:bool is_follow
}

struct BaseResp {
    1:i64 status_code
    2:string status_message
}

struct CreateUserReq{
    1:string username
    2:string password
}

struct CreateUserResp{
    1:BaseResp base_resp
    2:i64 user_id
}

struct GetUserByIdReq{
    1:i64 id
}

struct GetUserByIdResp{
    1:BaseResp base_resp
    2:User user
}

struct CheckUserReq {
    1: string username
    2: string password
}

struct CheckUserResp {
    1: i64 user_id
    2: BaseResp base_resp
}

service UserService {
    CreateUserResp CreateUser(1:CreateUserReq req)
    GetUserByIdResp GetUserById(1:GetUserByIdReq req)
    CheckUserResp CheckUser(1: CheckUserReq req)
}
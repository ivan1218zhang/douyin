namespace go message

include 'common.thrift'

struct MessageChatReq{
    1:i64 user_id
    2:i64 from_user_id
}
struct Message{
    1:i64 id
    2:i64 to_user_id
    3:i64 from_user_id
    4:string content
    5:i64 create_time
}

struct MessageChatResp{
    1:common.BaseResp base_resp
    2:list<Message> massage_list
}

struct MassageActionReq {
    1:i64 user_id
    2:i64 to_user_id
    3:i32 action_type
    4:string content
}

struct MassageActionResp{
    1:common.BaseResp base_resp
}

service FavoriteService {
    MessageChatResp GetMassageChat(1:MessageChatReq req)
    MassageActionResp MassageAction(1:MassageActionReq req)
}

namespace go comment

include 'common.thrift'

struct CommentActionReq{
    1:i64 user_id
    2:i64 video_id
    3:i32 action_type
    4:string comment_text
    5:i64 comment_id
}

struct CommentActionResp{
    1:common.BaseResp base_resp
    2:common.Comment comment
}

struct MGetCommentReq{
    1:i64 video_id
}

struct MGetCommentResp{
    1:common.BaseResp base_resp
    2:list<common.Comment> comment_list
}

struct CountCommentReq{
    1:i64 video_id
}

struct CountCommentResp{
    1:common.BaseResp base_resp
    2:i64 comment_count
}

struct MCountCommentReq{
    1:list<i64> video_id_list
}

struct MCountCommentResp{
    1:common.BaseResp base_resp
    2:list<i64> comment_count_list
}

service CommentService {
    CommentActionResp CommentAction(1:CommentActionReq req)
    CountCommentResp CountComment(1:CountCommentReq req)
    MCountCommentResp MCountComment(1:MCountCommentReq req)
    MGetCommentResp MGetComment(1:MGetCommentReq req)
}
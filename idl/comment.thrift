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

struct GetCommentsByVideoIdReq{
    1:i64 video_id
}

struct GetCommentsByVideoIdResp{
    1:common.BaseResp base_resp
    2:list<common.Comment> comments
}

struct CountCommentReq{
    1:i64 video_id
}

struct CountCommentResp{
    1:common.BaseResp base_resp
    2:i64 comment_count
}

struct CountCommentListReq{
    1:list<i64> video_id
}

struct CountCommentListResp{
    1:common.BaseResp base_resp
    2:list<i64> comment_count
}

service CommentService {
    CommentActionResp CreateComment(1:CommentActionReq req)
    CountCommentResp CountComment(1:CountCommentReq req)
    CountCommentListResp CountCommentList(1:CountCommentListReq req)
    GetCommentsByVideoIdResp GetCommentsByVideoId(1:GetCommentsByVideoIdReq getCommentsByVideoIdReq)
}
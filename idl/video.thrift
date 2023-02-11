include 'common.thrift'

namespace go video

struct MGetVideoReq {
    1:i64 latest_time
    2:i64 user_id
}

struct MGetVideoResp {
    1:common.BaseResp base_resp
    2:list<common.Video> video_list
    3:i64 next_time
}

struct PublishReq{
    1:i64 user_id
    2:string title
    3:binary data
}

struct PublishResp{
    1:common.BaseResp base_resp
}

struct MGetPublishReq{
    1:i64 user_id
}

struct MGetPublishResp{
    1:common.BaseResp base_resp
    2:list<common.Video> video_list
}

service VideoService {
    MGetVideoResp MGetVideo(1:MGetVideoReq req)
    PublishResp Publish(1:PublishReq req)
    MGetPublishResp MGetPublish(1:MGetPublishReq req)
}
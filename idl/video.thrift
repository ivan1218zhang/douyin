include 'common.thrift'

namespace go video

struct GetVideosReq {
    1:i64 latest_time
    2:i64 user_id
}

struct GetVideosResp {
    1:common.BaseResp base_resp
    2:list<common.Video> videos
    3:i64 nextTime
}

struct PublishReq{
    1:i64 user_id
    2:string title
    3:list<byte> data
}

struct PublishResp{
    1:common.BaseResp base_resp
    2:string url
}

struct GetPublishListReq{
    1:i64 user_id
}

struct GetPublishListResp{
    1:common.BaseResp base_resp
    2:list<common.Video> videos
}

service VideoService {
    GetVideosResp GetVideos(1:GetVideosReq getVideosReq)
    PublishResp Publish(1:PublishReq publishReq)
    GetPublishListResp GetPublishList(1:GetPublishListReq getPublishListReq)
}
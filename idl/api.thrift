namespace go api

include 'common.thrift'

struct MGetVideoReq {
    1:i64 latest_time
    2:string token
}

struct MGetVideoResp {
    1:i64 status_code
    2:string status_message
    3:list<common.Video> video_list
    4:i64 next_time
}

service ApiService {
    MGetVideoResp MGetVideo(1:MGetVideoReq req) (api.get="/douyin/feed/")
}
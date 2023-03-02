namespace go videodouyin

include "base.thrift"

struct FeedRequest {
    1:optional i64 latest_time // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
    2:optional i64 user_id // 可选参数，登录用户设置
}

struct FeedResponse {
    1:i32 status_code 
    2:optional string status_msg
    3:list<base.Video> video_list // 视频列表
    4:optional i64 next_time // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

struct PublishActionRequest {
    1:string token
    2:binary data // 视频数据
    3:string title // 视频标题
}

struct PublishActionResponse {
    1:i32 status_code
    2:optional string status_msg
}

service VideoService {
    FeedResponse Feed(1:FeedRequest req)
    PublishActionResponse PublishAction(1:PublishActionRequest req)
}
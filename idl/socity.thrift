namespace go socitydouyin

include "base.thrift"

struct FavoriteActionRequest {
    1:i64 user_id // 用户id
    2:i64 video_id // 视频id
    3:i32 action_type // 1-点赞，2-取消点赞
}

struct FavoriteActionResponse {
    1:i32 status_code 
    2:optional string status_msg
}

struct CommentActionRequest {
    1:i64 user_id
    2:i64 video_id
    3:i32 action_type // 1-发布评论，2-删除评论
    4:optional string comment_text // 用户填写的评论内容，在action_type=1的时候使用
    5:optional i64 comment_id // 要删除的评论id，在action_type=2的时候使用
}

struct CommentActionResponse {
    1:i32 status_code 
    2:optional string status_msg
    3:optional base.Comment comment // 评论成功返回评论内容，不需要重新拉取整个列表
}

struct CommentListRequest {
    1:i64 user_id
    2:i64 video_id
}

struct CommentListResponse {
    1:i32 status_code 
    2:optional string status_msg
    3:optional list<base.Comment> comment_list
}

service SocityService {
    FavoriteActionResponse FavoriteAction(1:FavoriteActionRequest req)
    CommentActionResponse CommentAction(1:CommentActionRequest req)
    CommentListResponse CommentList(1:CommentListRequest req)
}
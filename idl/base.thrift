struct User {
    1:i64 id // 用户id
    2:string name // 用户名称
    3:optional i64 follow_count // 关注总数
    4:optional i64 follower_count // 粉丝总数
    5:bool is_follow // true-已关注，false-未关注
    6:optional string avatar //用户头像
    7:optional string background_image //用户个人页顶部大图
    8:optional string signature //个人简介
    9:optional i64 total_favorited //获赞数量
    10:optional i64 work_count //作品数量
    11:optional i64 favorite_count //点赞数量
}

struct Video {
    1:i64 id // 视频唯一标识
    2:User author // 视频作者信息
    3:string play_url // 视频播放地址
    4:string cover_url // 视频封面地址
    5:i64 favorite_count // 视频点赞总数
    6:i64 comment_count // 视频评论总数
    7:bool is_favorite // true-已点赞，false-未点赞
    8:string title // 视频标题
    9:i64 create_time // 创建时间
}

struct Comment {
    1:i64 id
    2:User user
    3:string content
    4:string create_date // 评论发布日期,mm-dd
}
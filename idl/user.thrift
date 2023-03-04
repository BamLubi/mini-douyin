namespace go userdouyin

include "base.thrift"

struct UserRegisterRequest {
    1:string username // 注册用户名，最长32个字符
    2:string password // 密码，最长32个字符
}

struct UserRegisterResponse {
    1:i32 status_code
    2:optional string status_msg
    3:i64 user_id
    4:string token
}

struct UserLoginRequest {
    1:string username
    2:string password
}

struct UserLoginResponse {
    1:i32 status_code
    2:optional string status_msg
    3:i64 user_id
    4:string token
}

struct UserInfoRequest {
    1:i64 user_id
}

struct UserInfoResponse {
    1:i32 status_code
    2:optional string status_msg
    3:base.User user
}

struct PublishListRequest {
    1:i64 user_id
}

struct PublishListResponse {
    1:i32 status_code
    2:optional string status_msg
    3:list<base.Video> video_list // 用户发布的视频列表
}

struct FavoriteListRequest {
    1:i64 user_id
}

struct FavoriteListResponse {
    1:i32 status_code
    2:optional string status_msg
    3:list<base.Video> video_list
}

service UserService {
    UserRegisterResponse UserRegister(1:UserRegisterRequest req)
    UserLoginResponse UserLogin(1:UserLoginRequest req)
    UserInfoResponse UserInfo(1:UserInfoRequest req)
    PublishListResponse PublishList(1:PublishListRequest req)
    FavoriteListResponse FavoriteList(1:FavoriteListRequest req)
}
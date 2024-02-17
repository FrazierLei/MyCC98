namespace go user

struct User {
    1: i64 id               // 用户 ID
    2: string email         // 邮箱
    3: string nickname      // 昵称
    4: string password      // 密码
    5: string phone         // 手机号
    6: string about_me      // 自我介绍
    7: i64 post_cnt         // 发帖数
    8: i64 fan_cnt          // 粉丝数
    9: i64 follow_cnt       // 关注数
    10: string level        // 等级
    11: i64 wealth          // 财富值
}

struct SignupRequest {
    1: User user
}

struct SignupResponse {
}

struct LoginRequest {
    1: string email
    2: string password
}

struct LoginResponse {
    1: User user
}

struct FindOrCreateRequest {
    1: string phone
}

struct FindOrCreateResponse {
    1: User user
}

struct ProfileRequest {
    1: i64 id
}

struct ProfileResponse {
    1: User user
}

struct UpdateNonSensitiveInfoRequest {
    1: User user
}

struct UpdateNonSensitiveInfoResponse {

}


service UserService {
    SignupResponse Signup(1: SignupRequest req)
    LoginResponse Login(1: LoginRequest req)
    FindOrCreateResponse FindOrCreate(1: FindOrCreateRequest req)
    ProfileResponse Profile(1: ProfileRequest req)
    UpdateNonSensitiveInfoResponse UpdateNonSensitiveInfo(1: UpdateNonSensitiveInfoRequest req)
}
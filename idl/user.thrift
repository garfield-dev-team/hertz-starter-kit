namespace go user

struct BaseResp {
    1: i64 statusCode;
    2: string statusMsg;
    3: string data;
}

struct RegisterRequest {
    1: string username;
    2: string password;
}

struct RegisterResponse {
    1: BaseResp baseResp;
}

struct LoginRequest {
    1: string username;
    2: string password;
}

struct LoginResponse {
    1: BaseResp baseResp;
}

struct InfoRequest {
    1: string username;
}

struct InfoResponse {
    1: BaseResp baseResp;
}

service UserService {
    RegisterResponse Register(1: RegisterRequest req) (api.post="/user/register");
    LoginResponse Login(1: LoginRequest req) (api.post="/user/login");
    InfoResponse Info(1: InfoRequest req) (api.get="/user/:username");
}

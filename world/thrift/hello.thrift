namespace go hellothrift

typedef map<string, string> Data

struct Error {
    
}

struct ReqHeader {
    1: required string trace_id;
}

struct RespHeader {
    1: required string trace_id;
    2: required i32 code;
    3: required string msg;
}

struct Response {
    1:required i32 code; //错误码
    2:required string message; //错误信息
    3:required Data data;
}


struct HelloReq{
    1:required ReqHeader header;
    2:required string name;
}

struct HelloResp{
    1:required RespHeader header;
    2:required Data data;
}

service HelloService{
    HelloResp SayHello(1:HelloReq req)
}
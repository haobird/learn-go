include "User.thrift"

namespace go Sample
namespace php Sample
namespace java Sample

typedef map<string, string> Data

exception BizException {
    1:required i32 code
    2:required string msg
}

struct Response {
    1:required i32 errCode; //错误码
    2:required string errMsg; //错误信息
    3:required Data data;
}

//定义服务
service Greeter {
    Response SayHello(1:required User.User user) throws (1: BizException e)
    Response GetUser(1:required i32 uid) throws (1: BizException e)
}

namespace go Sample
namespace php Sample
namespace java Sample

struct User {
    1:required i32 id;
    2:required string name;
    3:required string avatar;
    4:required string address;
    5:required string mobile;
}

struct UserList {
    1:required list<User> userList;
    2:required i32 page;
    3:required i32 limit;
}


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
    Response SayHello(1:required User user) throws (1: BizException e)
    Response GetUser(1:required i32 uid) throws (1: BizException e)
}

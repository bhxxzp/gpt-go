// idl/hello.thrift
namespace go hello

struct HelloReq {
    1: required string Name (api.query="name"); // 添加 api 注解为方便进行参数绑定
}

struct HelloResp {
    1: string RespBody;
}


service HelloService {
    HelloResp Hello(1: HelloReq request) (api.get="/hello");
}

struct ShitReq {
    1: string Name (api.query="name"); // 添加 api 注解为方便进行参数绑定
}

struct ShitResp {
    1: string RespBody;
}


service ShitService {
    HelloResp Shit(1: HelloReq request) (api.get="/hello/shit");
}

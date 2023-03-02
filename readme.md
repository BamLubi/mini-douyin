

# Kitex RPC 框架使用

## 快速开始
[Kitex官网快速开始](https://www.cloudwego.io/zh/docs/kitex/getting-started/)

[Kitex代码生成](https://www.cloudwego.io/zh/docs/kitex/tutorials/code-gen/code_generation/)


1. 安装 kitex 和 thriftgo

```shell
go install github.com/cloudwego/kitex/tool/cmd/kitex@latest
go install github.com/cloudwego/thriftgo@latest

# 升级框架
go get github.com/cloudwego/kitex@latest
go mod tidy
```

2. 生成服务端代码文件

```shell
kitex -module "your_module_name" -service "your_service_name" hello.thrift
# userservice
cd cmd/user
kitex -module mini-douyin -service userservice ../../idl/user.thrift
# videoservice
cd cmd/video
kitex -module mini-douyin -service videoservice ../../idl/video.thrift

```

3. 生成客户端代码文件

```shell
cd cmd/api
kitex -module mini-douyin ../../idl/user.thrift
kitex -module mini-douyin ../../idl/video.thrift
```

4. 编译运行

```shell
sh build.sh
sh output/bootstrap.sh
```


jwt鉴权组件 https://github.com/appleboy/gin-jwt

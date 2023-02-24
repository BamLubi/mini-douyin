

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
kitex -service "your_service_name" hello.thrift
# 若当前目录不在 $GOPATH/src 下，需要加上 -module 参数，一般为 go.mod 下的名字
kitex -module "your_module_name" -service "your_service_name" hello.thrift
```

3. 生成客户端代码文件

```shell
kitex hello.thrift
```

4. 编译运行

```shell
sh build.sh
sh output/bootstrap.sh
```

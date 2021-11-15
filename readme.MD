
# GRPC

### mac安装 protoc 编译器
* https://github.com/google/protobuf/releases
* 下载对应系统版本
* 放到目录
* /usr/local/bin 


### 使用以下命令为 Go 安装协议编译器插件：
```shell
go get -u github.com/golang/protobuf/protoc-gen-go
# 或者
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
```
```shell

option go_package="/services"; // 指定包名 生成的go文件包名
进入 pbfiles 文件夹 执行以下命令
protoc --go_out=../ Prod.proto
```
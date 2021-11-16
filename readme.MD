
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



[comment]: <> (在 proto 文件中定义生成的go文件包名)

[comment]: <> (```shell)

[comment]: <> (option go_package="/services"; // 指定包名 )

[comment]: <> (```)

[comment]: <> (进入 pbfiles 文件夹 执行以下命令)

[comment]: <> (```shell)

[comment]: <> (protoc --go_out=../ Prod.proto)

[comment]: <> (```)

[comment]: <> (将会在`services`目录下生成`Prod.pb.go`文件)

[comment]: <> (在proto文件中增加)

[comment]: <> (```protobuf)

[comment]: <> (service ProdService{)

[comment]: <> ( rpc GetProdStock &#40;ProdRequest&#41; returns &#40;ProdResponse&#41;;)

[comment]: <> (})

[comment]: <> (```)

[comment]: <> (执行以下命令生成文件 Prod_grpc.pb.go)

[comment]: <> (```shell)

[comment]: <> (protoc --go-grpc_out=../ Prod.proto)

[comment]: <> (```)

[comment]: <> (关于生成的这些东西有点乱 命令也是乱的不行)

[comment]: <> (更详细的资料看这里)

[comment]: <> (https://developers.google.com/protocol-buffers)

[comment]: <> (https://github.com/protocolbuffers/protobuf-go/releases/tag/v1.20.0#v1.20-generated-code)



未完待续。。。
package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	hello_grpc "grpc/pb"
	"net"
)

// 服务端
// 1、取出server
// 2、挂载方法
// 3、注册服务
// 4、创建监听

// 1、取出server

type server struct {
	hello_grpc.UnimplementedHelloGRPCServer
}

// 2、挂载方法

func (s *server) SayHi(ctx context.Context, req *hello_grpc.Req) (res *hello_grpc.Res, err error) {
	fmt.Println(req.GetMessage())
	return &hello_grpc.Res{Message: "我是服务端返回的grpc内容"}, nil
}

func main() {

	// 3、注册服务
	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println(err.Error())
	}
	s := grpc.NewServer()
	hello_grpc.RegisterHelloGRPCServer(s, &server{})

	// 4、创建监听
	err = s.Serve(lis)
	if err != nil {
		fmt.Println(err.Error())
	}
}

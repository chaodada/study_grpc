package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	hello_grpc "grpc/pb"
)

// 客户端
// 1、创建一个链接
// 2、new一个client
// 3、调用client的方法
// 4、获取返回值

func main() {
	// 1、创建一个链接
	conn, err := grpc.Dial("localhost:8888",grpc.WithInsecure())
	if err != nil {
		fmt.Println(err.Error())
	}
	defer conn.Close()

	// 2、new一个client
	client := hello_grpc.NewHelloGRPCClient(conn)

	// 3、调用client的方法
	res, err := client.SayHi(context.Background(), &hello_grpc.Req{Message: "我是客户端发送到服务端的数据"})
	if err != nil {
		fmt.Println(err.Error())
	}

	// 4、获取返回值
	fmt.Println(res.GetMessage())

}

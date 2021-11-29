package main

import (
	"context"
	person "demo2/pb/person"
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	// 1、创建一个链接
	conn, err := grpc.Dial("localhost:8888",grpc.WithInsecure())
	if err != nil {
		fmt.Println(err.Error())
	}
	defer conn.Close()

	// 2、new一个client
	client := person.NewSearchServiceClient(conn)

	// 3、调用client的方法
	// 传统的 即刻响应
	res, err := client.Search(context.Background(), &person.PersonReq{Name: "名字",Age: 18})
	if err != nil {
		fmt.Println(err.Error())
	}

	// 4、获取返回值
	fmt.Printf("%#v\n",res.Name)
	fmt.Println(res)
}

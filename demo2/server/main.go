package main

import (
	"context"
	"demo2/pb/person"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

// 1、取出server

type personServer struct {
	person.UnimplementedSearchServiceServer
}

// 2、挂载方法

// Search 传统的 即刻响应
func (p *personServer) Search(ctx context.Context, req *person.PersonReq) (*person.PersonRes, error) {
	name := req.GetName()
	res := &person.PersonRes{Name: "服务器收到了" + name + "的信息"}
	return res, nil
}

func (p *personServer) SearchIn(person.SearchService_SearchInServer) error {
	return nil
}
func (p *personServer) SearchOut(*person.PersonReq, person.SearchService_SearchOutServer) error {
	return nil
}
func (p *personServer) SearchIO(person.SearchService_SearchIOServer) error {
	return nil
}

func main() {
	// 3、注册服务
	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println(err.Error())
	}
	s := grpc.NewServer()
	person.RegisterSearchServiceServer(s, &personServer{})

	// 4、创建监听
	err = s.Serve(lis)
	if err != nil {
		fmt.Println(err.Error())
	}

}

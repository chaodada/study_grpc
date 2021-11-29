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

// SearchIn 入参为流的形式   客户端发送流  服务端接收流  并且服务端会给客户端一个流式反馈 【在停止接收的时候】
func (p *personServer) SearchIn(server person.SearchService_SearchInServer) error {
	// server.SendAndClose(*PersonRes) error 流式反馈
	// server.Recv() (*PersonReq, error) 不断的读取req 过来流就读一下过来流就读一下

	for {
		req, err := server.Recv()
		fmt.Println(req)
		if err != nil {
			server.SendAndClose(&person.PersonRes{Name: "服务端接收流完成"})
			break
		}

	}

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

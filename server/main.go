package main

import (
	pb "GRPC/server/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	pb.UnimplementedSayHelloServer
}

func (s *server) Say(ctx context.Context, req *pb.SayRequest) (*pb.SayResponse, error) {
	return &pb.SayResponse{Message: "hello" + req.GetName()}, nil
}

func main() {
	// 开启端口
	listen, _ := net.Listen("tcp", ":9090")
	// 创建grpc服务
	grpcServer := grpc.NewServer()
	// 在grpc服务端口中注册自己编写的服务
	pb.RegisterSayHelloServer(grpcServer, &server{})

	//启动服务
	err := grpcServer.Serve(listen)
	if err != nil {
		fmt.Printf("failed to serve : %v", err)
		return
	}
}

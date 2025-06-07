package main

import (
	pb "GRPC/server/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	// 连接server
	connect, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect %v", err)
	}
	defer connect.Close()

	// 建立连接
	client := pb.NewSayHelloClient(connect)

	// 执行rpc调用
	say, _ := client.Say(context.Background(), &pb.SayRequest{Name: "LC"})
	fmt.Println(say)
}

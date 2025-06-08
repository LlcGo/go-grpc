package main

import (
	pb "GRPC/server/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type PerRPCCredentials interface {
	GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error)
	RequireTransportSecurity() bool
}

type Client struct {
}

func (c Client) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appId":  "lc",
		"appKey": "666",
	}, nil
}

func (c Client) RequireTransportSecurity() bool {
	return false
}

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithPerRPCCredentials(new(Client)))
	// 连接server
	connect, err := grpc.NewClient("127.0.0.1:9090", opts...)
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

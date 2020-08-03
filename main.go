package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb_son "github.com/alx696/go-grpc-test/son"
)

const (
	port = ":40000"
)

type server struct{
	pb_son.UnimplementedSonServiceServer
}

//实现Insert接口
func (s *server) Insert(ctx context.Context, in *pb_son.Info) (*pb_son.Info, error) {
	log.Println("增加:", in.Name)
	return &pb_son.Info{Name: "你好" + in.Name}, nil
}

func main() {
	log.Println("服务启动")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	//注册服务
	pb_son.RegisterSonServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

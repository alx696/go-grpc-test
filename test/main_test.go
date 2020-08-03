package test

import (
	"context"
	"log"
	"testing"
	"time"

	"google.golang.org/grpc"
	pb_son "github.com/alx696/go-grpc-test/son"
)

// TestTime 测试时间
func TestTime(t *testing.T) {
	log.Println(time.Now().Format("2006-01-02 15:04:05-07"))
}

func TestGrpc(t *testing.T) {
	conn, err := grpc.Dial("localhost:40000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb_son.NewSonServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Insert(ctx, &pb_son.Info{Name: "小明"})
	if err != nil {
		log.Fatalf("调用失败: %v", err)
	}
	log.Println("回复: ", r.GetName())
}

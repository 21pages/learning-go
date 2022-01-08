package main

import (
	"context"
	"log"
	"time"

	pb "helloworld/helloworld"

	"google.golang.org/grpc"
)

func main() {
	//建立连接
	conn, err := grpc.Dial("localhost:50000", grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect", err)
		return
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	//request and get response
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	rsp, err := c.SayHello(ctx, &pb.HelloRequest{Name: "zhangsan"})
	if err != nil {
		log.Fatal("request failed:", err)
		return
	}
	log.Println("response:", rsp.GetMessage())
}

/*
output:
2022/01/08 14:59:05 response: Hello:zhangsan
*/

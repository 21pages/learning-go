package main

import (
	"context"
	"fmt"
	pb "helloworld/helloworld"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

//SayHello implement "SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)"
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Println("Received:", in.GetName())
	return &pb.HelloResponse{Message: "Hello:" + in.GetName()}, nil
}

func main() {
	listen, err := net.Listen("tcp", fmt.Sprintf("localhost:50000"))
	if err != nil {
		log.Fatalln("Failed to listen!")
		return
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{}) //参数是interface{}, 要传结构体指针
	log.Println("server listen at", listen.Addr())
	if err := s.Serve(listen); err != nil { //Serve 内部有for
		log.Fatalln("failed to serve:", err)
	}
}

/*
output:
2022/01/08 14:59:02 server listen at 127.0.0.1:50000
2022/01/08 14:59:05 Received: zhangsan
*/

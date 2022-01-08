package main

import (
	"context"
	"fmt"
	pb "fourmode/four"
	"io"
	"log"
	"net"
	"strconv"
	"strings"

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) NoStream(ctx context.Context, req *pb.ReqMsg) (*pb.RspMsg, error) {
	return &pb.RspMsg{Reply: fmt.Sprintf("reply:%v", req.GetRequest())}, nil
}

func (s *server) ClientStream(srv pb.Four_ClientStreamServer) error {
	replyArr := []string{}
	for {
		req, err := srv.Recv()
		if err == nil {
			replyArr = append(replyArr, req.GetRequest())
		} else if err == io.EOF {
			reply := strings.Join(replyArr, ";")
			if err := srv.SendAndClose(&pb.RspMsg{Reply: reply}); err != nil {
				log.Fatalln("ClientStream send failed:", err)
				return err
			}
			return nil
		} else {
			log.Fatalln("ClientStream Recv:", err)
			return err
		}
	}
}

func (s *server) ServerStream(req *pb.ReqMsg, srv pb.Four_ServerStreamServer) error {
	for i := 0; i < 10; i++ {
		if err := srv.Send(&pb.RspMsg{Reply: fmt.Sprintf("%s-%s", req.GetRequest(), strconv.Itoa(i))}); err != nil {
			log.Fatal("ServerStream Send:", err)
			return err
		}
	}
	return nil
}

func (s *server) DoubleStream(srv pb.Four_DoubleStreamServer) error {
	for {
		req, err := srv.Recv()
		if err == io.EOF {
			return nil
		} else if err != nil {
			log.Fatalln("DoubleStream Recv:", err)
			return err
		} else {
			srv.Send(&pb.RspMsg{Reply: req.GetRequest()})
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:50000")
	if err != nil {
		log.Fatalln("listen:", err)
	}
	log.Println("server start!")
	s := grpc.NewServer()
	pb.RegisterFourServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalln("Serve:", err)
	}
}

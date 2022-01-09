package main

import (
	"context"
	"log"
	"net"
	pb "pubsub/proto"
	"strings"
	"time"

	"github.com/docker/docker/pkg/pubsub"
	"google.golang.org/grpc"
)

type PubsubServiceServer struct {
	pub *pubsub.Publisher //docker的pubsub
}

func NewPubsubServiceServer() *PubsubServiceServer {
	return &PubsubServiceServer{
		pub: pubsub.NewPublisher(100*time.Millisecond, 10),
	}
}

func (s *PubsubServiceServer) Publish(ctx context.Context, arg *pb.StringPub) (*pb.StringPub, error) {
	s.pub.Publish(arg.GetValue())
	return &pb.StringPub{}, nil
}

func (s *PubsubServiceServer) Subscribe(arg *pb.StringPub, stream pb.PubsubService_SubscribeServer) error {
	ch := s.pub.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, arg.GetValue()) {
				return true
			}
		}
		return false
	})

	for v := range ch {
		if err := stream.Send(&pb.StringPub{Value: v.(string)}); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPubsubServiceServer(s, NewPubsubServiceServer())
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

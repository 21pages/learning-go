package main

import (
	"context"
	"fmt"
	"io"
	"log"
	pb "pubsub/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewPubsubServiceClient(conn)

	stream, err := client.Subscribe(context.Background(), &pb.StringPub{Value: "golang:"})
	if err != nil {
		log.Fatal(err)
	}

	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		fmt.Println(reply.GetValue())
	}
}

/*
output:
golang:hello golang
*/

package main

import (
	"context"
	"log"

	pb "pubsub/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":50000", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Dial:", err)
	}
	defer conn.Close()

	cli := pb.NewPubsubServiceClient(conn)

	_, err = cli.Publish(context.Background(), &pb.StringPub{Value: "golang:hello golang"})
	if err != nil {
		log.Fatalln("Publish:", err)
	}

	_, err = cli.Publish(context.Background(), &pb.StringPub{Value: "docker:hello docker"})
	if err != nil {
		log.Fatalln("Publish:", err)
	}

}

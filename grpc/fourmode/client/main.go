package main

import (
	"context"
	"fmt"
	pb "fourmode/four"
	"io"
	"log"
	"strings"

	"google.golang.org/grpc"
)

func main() {
	//如果不设置WithInsecure:Dial: grpc: no transport security set (use grpc.WithInsecure() explicitly or set credentials)
	conn, err := grpc.Dial("localhost:50000", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Dial:", err)
	}
	client := pb.NewFourClient(conn)

	//nostream
	rsp, err := client.NoStream(context.Background(), &pb.ReqMsg{Request: "no-stream-req-msg"}) //直接发
	if err != nil {
		log.Fatalln("nostream, err:", err)
	}
	log.Println("nostream, reply:", rsp.GetReply())
	//cancel1()

	//clientstream
	clientStreamClient, err := client.ClientStream(context.Background()) //获取client, 而不是直接发
	if err != nil {
		log.Fatalln("ClientStream, err:", err)
	}
	for i := 0; i < 10; i++ {
		if err := clientStreamClient.Send(&pb.ReqMsg{Request: fmt.Sprintf("clientStream%d", i)}); err != nil {
			log.Fatal("ClientStream:", i, err)
		}
	}
	rsp, err = clientStreamClient.CloseAndRecv()
	if err != nil {
		log.Fatal("clientStream Recv Failed:", err)
	}
	log.Println("clientStream Recv:", rsp.GetReply())
	//cancel2()

	//serverstream
	serverStreamClient, err := client.ServerStream(context.Background(), &pb.ReqMsg{Request: "server-stream-msg"}) //直接发
	if err != nil {
		log.Fatalln("serverstream, err:", err)
	}
	for {
		rsp, err := serverStreamClient.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal("serverstream Recv:", err)
		}
		log.Println("serverstream recv rsp:", rsp.GetReply())
	}
	//cancel3()

	//doublesteram
	doubleStreamClient, err := client.DoubleStream(context.Background())
	if err != nil {
		log.Fatalln("doublesteram, err:", err)
	}
	waitchan := make(chan struct{})
	go func() {
		var cnt int
		for {
			rsp, err := doubleStreamClient.Recv()
			if err != nil && err != io.EOF {
				log.Println("doublesteram Recv:", err)
				continue
			}
			reply := rsp.GetReply()
			log.Println("doublesteram recv rsp:", reply)
			if strings.Contains(reply, "end") {
				close(waitchan)
				break
			}
			cnt++
			if cnt < 10 {
				doubleStreamClient.Send(&pb.ReqMsg{Request: fmt.Sprintf("double-stream-msg-%d", cnt)})
			} else {
				doubleStreamClient.Send(&pb.ReqMsg{Request: "end"})
				doubleStreamClient.CloseSend()
			}
		}
	}()
	doubleStreamClient.Send(&pb.ReqMsg{Request: "start"})
	<-waitchan
}

/*
output:
2022/01/08 21:44:50 nostream, reply: reply:no-stream-req-msg
2022/01/08 21:44:50 clientStream Recv: clientStream0;clientStream1;clientStream2;clientStream3;clientStream4;clientStream5;clientStream6;clientStream7;clientStream8;clientStream9
2022/01/08 21:44:50 serverstream recv rsp: server-stream-msg-0
2022/01/08 21:44:50 serverstream recv rsp: server-stream-msg-1
2022/01/08 21:44:50 serverstream recv rsp: server-stream-msg-2
2022/01/08 21:44:50 serverstream recv rsp: server-stream-msg-3
2022/01/08 21:44:50 serverstream recv rsp: server-stream-msg-4
2022/01/08 21:44:50 serverstream recv rsp: server-stream-msg-5
2022/01/08 21:44:50 serverstream recv rsp: server-stream-msg-6
2022/01/08 21:44:50 serverstream recv rsp: server-stream-msg-7
2022/01/08 21:44:50 serverstream recv rsp: server-stream-msg-8
2022/01/08 21:44:50 serverstream recv rsp: server-stream-msg-9
2022/01/08 21:44:50 doublesteram recv rsp: start
2022/01/08 21:44:50 doublesteram recv rsp: double-stream-msg-1
2022/01/08 21:44:50 doublesteram recv rsp: double-stream-msg-2
2022/01/08 21:44:50 doublesteram recv rsp: double-stream-msg-3
2022/01/08 21:44:50 doublesteram recv rsp: double-stream-msg-4
2022/01/08 21:44:50 doublesteram recv rsp: double-stream-msg-5
2022/01/08 21:44:50 doublesteram recv rsp: double-stream-msg-6
2022/01/08 21:44:50 doublesteram recv rsp: double-stream-msg-7
2022/01/08 21:44:50 doublesteram recv rsp: double-stream-msg-8
2022/01/08 21:44:50 doublesteram recv rsp: double-stream-msg-9
2022/01/08 21:44:50 doublesteram recv rsp: end
*/

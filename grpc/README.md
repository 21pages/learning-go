# rpc的四种类型

rpc需要在proto文件中定义  
```go
service MyService{
    rpc ServiceFunction(MessageClient) returns (MessageServer) {}
    rpc ServiceFunction(stream MessageClient) returns (MessageServer) {}
    rpc ServiceFunction(MessageClient) returns (stream MessageServer) {}
    rpc ServiceFunction(stream MessageClient) returns (stream MessageServer) {}
}
```


1. 无stream:  客户端发送后等待服务器返回

2. 客户端stream: 客户端流式发送, 全部发送完毕后,服务器开始响应

3. 服务器stream: 客户端发送后等待服务器流式响应

4. 双stream:客户端可以一直请求, 服务器也可以一直响应

# 安装
sudo apt-get install golang-goprotobuf-dev

# 使用方法

1. 编写proto文件, 定义Service, Message

2. protoc --go_out=plugins=grpc:. *.proto

3. 编写客户端和服务器

以helloworld为例  

helloworld.pb.go:

```go
// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloResponse, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld.proto",
}
```

pb默认实现UnimplementedGreeterServer, 自行实现:
```go
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Println("Received:", in.GetName())
	return &pb.HelloResponse{Message: "Hello:" + in.GetName()}, nil
}
```

main:

```go
func main() {
	listen, err := net.Listen("tcp", fmt.Sprintf("localhost:50000"))
	if err != nil {
		log.Fatalln("Failed to listen!")
		return
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{}) //参数是interface{}, 要传结构体指针
    									//未实现时传入&pb.UnimplementedGreeterServer{}
	log.Println("server listen at", listen.Addr())
	if err := s.Serve(listen); err != nil { //Serve 内部有for
		log.Fatalln("failed to serve:", err)
	}
}
```

client:

```go
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
```




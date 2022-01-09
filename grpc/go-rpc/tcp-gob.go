package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type HelloService struct{}

/*
1. 方法必须有两个可序列的参数, 第二个是指针
2. 返回值是error
3. public方法,首字母大写
*/
func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello " + request
	return nil
}

func main() {

	//server

	//注册服务
	rpc.RegisterName("HelloService", &HelloService{})

	//监听端口
	lis, err := net.Listen("tcp", ":50000")
	if err != nil {
		log.Fatalln("Listen:", err)
	}
	go func() {
		for {
			conn, err := lis.Accept()
			if err != nil {
				log.Fatalln("Accept:", err)
			}
			go rpc.ServeConn(conn) //默认gob
		}
	}()

	//client

	cli, err := rpc.Dial("tcp", ":50000")
	if err != nil {
		log.Fatalln("Dial:", err)
	}
	var reply string
	//服务名.方法名
	if err := cli.Call("HelloService.Hello", "buddy", &reply); err != nil {
		log.Fatalln("Call:", err)
	}
	fmt.Println("Client recv:", reply)

}

/*
output:
Client recv: hello buddy
*/

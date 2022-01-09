package main

import (
	"fmt"
	"log"
	"net/http"
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

	go func() {
		for {
			go rpc.HandleHTTP() //默认gob
			err := http.ListenAndServe(":50000", nil)
			if err != nil {
				log.Fatalln("ListenAndServe:", err)
			}
		}
	}()

	//client

	cli, err := rpc.DialHTTP("tcp", ":50000")
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

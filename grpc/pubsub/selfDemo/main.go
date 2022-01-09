package main

import (
	. "selfDemo/pubsub"
)

func main() {
	c1 := &Client{Id: 100, Ip: "172.18.1.1"}
	c3 := &Client{Id: 300, Ip: "172.18.1.3"}

	srv := NewServer()
	srv.Subscribe(c1, "Topic")
	srv.Subscribe(c3, "Topic")

	srv.PublishMessage("Topic", "测试信息1")

	srv.Unsubscribe(c3, "Topic")
	srv.PublishMessage("Topic", "测试信息2222")

	srv.Subscribe(c1, "Topic2")
	srv.Subscribe(c3, "Topic2")
	srv.PublishMessage("Topic2", " Topic2的测试信息")
}

/*
output:
channel:Topic client:300 message:测试信息1
channel:Topic client:100 message:测试信息1
channel:Topic client:100 message:测试信息2222
从channel:Topic 中删除client:300
channel:Topic2 client:300 message: Topic2的测试信息
channel:Topic2 client:100 message: Topic2的测试信息
*/

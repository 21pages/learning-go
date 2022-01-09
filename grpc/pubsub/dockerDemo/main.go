package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/docker/docker/pkg/pubsub"
)

func main() {
	pub := pubsub.NewPublisher(100*time.Microsecond, 10)
	//订阅golang:主题
	golangChan := pub.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, "golang:") {
				return true
			}
		}
		return false
	})
	//订阅docker:主题
	dockerChan := pub.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, "docker:") {
				return true
			}
		}
		return false
	})

	go pub.Publish("sun")
	go pub.Publish("golang: https://golang.org")
	go pub.Publish("docker: https://www.docker.com")

	time.Sleep(time.Second * 2)
	go func() {
		fmt.Println("golang topic:", <-golangChan)
	}()

	go func() {
		fmt.Println("docker topic:", <-dockerChan)
	}()
	time.Sleep(time.Second * 3)
	fmt.Println("end")
}

/*
output:
docker topic: docker: https://www.docker.com
golang topic: golang: https://golang.org
end
*/

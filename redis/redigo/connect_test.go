package redigo

import (
	"testing"

	"github.com/garyburd/redigo/redis"
)

func TestConnect(t *testing.T) {
	//默认端口6379
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		t.Fatal("Dial:", err)
	}
	defer conn.Close()
	t.Log("redis connect success")
}

/*
output:
2022/01/10 09:44:42 redis connect success
*/

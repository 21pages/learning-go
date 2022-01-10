package redigo

import (
	"testing"
	"time"

	"github.com/garyburd/redigo/redis"
)

func TestExpire(t *testing.T) {
	t.Cleanup(func() { conn.Close() })
	val := 123
	//expire前先set, 顺序颠倒不行
	if _, err := conn.Do("set", "abc", val); err != nil {
		t.Fatal(err)
	}
	//设置key="abc"的expire时间,秒
	if _, err := conn.Do("expire", "abc", 10); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 2; i++ {
		time.Sleep(time.Second * 6)
		reply, err := redis.Int(conn.Do("get", "abc"))
		if i == 0 {
			if err != nil {
				t.Fatal(err)
			}
			if val != reply {
				t.Fatal("reply:", reply)
			}
		} else {
			if err == nil {
				t.Fatal("can read")
			}
		}
	}

}

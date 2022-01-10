package redigo

import (
	"testing"
)

// 队列
func TestLpush(t *testing.T) {
	t.Cleanup(func() { conn.Close() })
	// val := []interface{}{"abc", "d", "e", "f"} //abc是key
	//设置key="abc"的expire时间,秒
	if _, err := conn.Do("lpush", "abc", "d", "e", "f", 10); err != nil {
		t.Fatal(err)
	}
	// val = val[1:]

	// for i := 0; i < 2; i++ {
	// 	reply, err := redis.Int(conn.Do("get", "abc"))
	// }

}

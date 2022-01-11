package redigo

import (
	"log"
	"os"
	"sort"
	"testing"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/stretchr/testify/assert"
)

var conn redis.Conn

func TestMain(m *testing.M) {
	//beforeTest -> init()

	code := m.Run()

	//afterTest
	conn.Close()

	os.Exit(code)
}

func init() {
	var err error
	conn, err = redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatalln("Dial:", err)
	}
	log.Println("redis connect success!")
}

func delKey(keys ...string) {
	for _, key := range keys {
		if _, err := conn.Do("del", key); err != nil {
			log.Fatal("del:", err)
		}
	}
}

func TestString(t *testing.T) {

	simple := func(t *testing.T) {
		key := "abc"
		val := "123"
		delKey(key)
		defer delKey(key)
		//set
		_, err := conn.Do("set", key, val)
		if err != nil {
			t.Fatal("Do set:", err)
		}

		//get
		tmp, _ := conn.Do("get", key)
		log.Printf("%T, %#v", tmp, tmp) //2022/01/12 11:55:15 []uint8, []byte{0x31, 0x32, 0x33}
		reply, err := redis.String(conn.Do("get", key))
		if err != nil {
			t.Fatal("String:", err)
		}

		assert.Equal(t, val, reply)
	}

	Int := func(t *testing.T) {
		key := "abc"
		val := 12

		delKey(key)
		defer delKey(key)
		//set
		_, err := conn.Do("set", key, val)
		if err != nil {
			t.Fatal("Do set:", err)
		}

		//get
		reply, err := redis.Int(conn.Do("get", key))
		if err != nil {
			t.Fatal("Int:", err)
		}

		assert.Equal(t, val, reply)
	}

	Ints := func(t *testing.T) {
		key := "abc"
		val := []int{1000, 2, 3}

		delKey(key)
		defer delKey(key)
		//set
		_, err := conn.Do("set", key, val)
		if err != nil {
			t.Fatal("Do set:", err)
		}
		//get
		/*
			//redigo: unexpected type for Values, got type []uint8
			reply, err := redis.Ints(conn.Do("get", key))
			if err != nil {
				t.Fatal("Ints:", err)
			}

			assert.Equal(t, val, reply)
		*/
		tmp, _ := conn.Do("get", key)
		log.Printf("%T, %#v\n", tmp, tmp) //int的内容实际上是字节流
		//2022/01/12 11:52:53 []uint8, []byte{0x5b, 0x31, 0x30, 0x30, 0x30, 0x20, 0x32, 0x20, 0x33, 0x5d}

		//unexpected type for Strings, got type []uint8
		// reply, err := redis.Strings(conn.Do("get", key))
		// if err != nil {
		// 	t.Fatal("Ints:", err)
		// }

		// assert.Equal(t, []string{"1000", "2", "3"}, reply)
	}

	IntMap := func(t *testing.T) {
		key := "abc"
		val := map[string]int{}
		val["a"] = 10000000
		val["b"] = 2

		delKey(key)
		defer delKey(key)
		//set
		_, err := conn.Do("set", key, val)
		if err != nil {
			t.Fatal("Do set:", err)
		}

		//get
		// reply, err := redis.IntMap(conn.Do("get", key))
		// if err != nil {
		// 	t.Fatal("IntMap:", err)
		// }

		// assert.Equal(t, val, reply)
	}

	mset := func(t *testing.T) {
		key1 := "hello"
		val1 := "HELLO"
		key2 := "world"
		val2 := "WORLD"

		delKey(key1, key2)
		defer delKey(key1, key2)
		//set
		_, err := conn.Do("mset", key1, val1, key2, val2)
		if err != nil {
			t.Fatal("Do set:", err)
		}

		//get
		reply, err := redis.Strings(conn.Do("mget", key1, key2))
		if err != nil {
			t.Fatal("String:", err)
		}
		sort.Strings(reply)
		vals := []string{val1, val2}
		sort.Strings(vals)

		assert.Equal(t, vals, reply)
	}

	expire := func(t *testing.T) {
		key := "abc"
		val := 123

		delKey(key)
		defer delKey(key)

		//expire前先set, 顺序颠倒不行
		if _, err := conn.Do("set", "abc", val); err != nil {
			t.Fatal(err)
		}
		//设置key="abc"的expire时间,秒
		if _, err := conn.Do("expire", "abc", 3); err != nil {
			t.Fatal(err)
		}

		for i := 0; i < 2; i++ {
			time.Sleep(time.Second * 2)
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

	tests := []struct {
		Name string
		test func(t *testing.T)
	}{
		{"simple", simple},
		{"Int", Int},
		{"Ints", Ints},
		{"IntMap", IntMap},
		{"mset", mset},
		{"expire", expire},
	}
	for _, f := range tests {
		t.Run(f.Name, f.test)
	}
}

func TestHash(t *testing.T) {
	key := "abc"
	fields := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fieldkeys := []string{}
	for k, _ := range fields {
		fieldkeys = append(fieldkeys, k)
	}
	sort.Strings(fieldkeys)
	for i := 0; i < len(fieldkeys); i++ {
		fieldkey := fieldkeys[i]
		if _, err := conn.Do("hset", key, fieldkey, fields[fieldkey]); err != nil {
			t.Fatal(err)
		}
	}
	reply, err := redis.IntMap(conn.Do("hgetall", key))
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, fields, reply)
}

func TestList(t *testing.T) {
	key := "abc"
	val := []string{"1", "2", "3"}

	delKey(key)
	defer delKey(key)

	for i := 0; i < len(val); i++ {
		if _, err := conn.Do("lpush", key, val[i]); err != nil {
			t.Fatal(err)
		}
	}

	for i := 0; i < len(val); i++ {
		reply, err := redis.String(conn.Do("rpop", key))
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, val[i], reply)
	}
}

func TestSet(t *testing.T) {
	key1 := "a"
	val1 := []string{"1", "2", "3"}
	key2 := "b"
	val2 := []string{"3", "4", "5"}

	delKey(key1, key2)
	defer delKey(key1, key2)

	for _, v := range val1 {
		if _, err := conn.Do("sadd", key1, v); err != nil {
			t.Fatal(err)
		}
	}
	for _, v := range val2 {
		if _, err := conn.Do("sadd", key2, v); err != nil {
			t.Fatal(err)
		}
	}

	reply, err := redis.Strings(conn.Do("smembers", key1))
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, val1, reply)

	tests := []struct {
		op     string
		expect []string
	}{
		{"sdiff", []string{"1", "2"}},
		{"sinter", []string{"3"}},
		{"sunion", []string{"1", "2", "3", "4", "5"}},
	}

	for i := 0; i < len(tests); i++ {
		reply, err = redis.Strings(conn.Do(tests[i].op, key1, key2))
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, tests[i].expect, reply)
	}
}

func TestSortedSet(t *testing.T) {
	key := "abc"
	members := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	delKey(key)
	defer delKey(key)

	for k, v := range members {
		if _, err := conn.Do("zadd", key, v, k); err != nil {
			t.Fatal(err)
		}
	}

	reply, err := redis.Strings(conn.Do("zpopmax", key))
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, reply, []string{"c", "3"})
}

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     16,
		MaxActive:   0, //auto
		IdleTimeout: time.Second * 300,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", ":6379")
		},
	}
}

func TestPool(t *testing.T) {
	c := pool.Get()
	if c == nil {
		t.Fatal("pool Get")
	}
	defer c.Close()

	key := "abc"
	val := "hello"

	delKey(key)
	defer delKey(key)

	_, err := c.Do("set", key, val)
	if err != nil {
		t.Fatal(err)
	}
	reply, err := redis.String(c.Do("get", key))
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, val, reply)
}

func BenchmarkGet(b *testing.B) {
	key := "abc"
	val := "hello"
	delKey(key)
	defer delKey(key)

	if _, err := conn.Do("set", key, val); err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reply, err := redis.String(conn.Do("get", key))
		if err != nil {
			b.Fatal(err)
		}
		assert.Equal(b, val, reply)
	}

}

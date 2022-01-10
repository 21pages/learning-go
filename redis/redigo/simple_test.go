package redigo

import (
	"log"
	"testing"

	"github.com/garyburd/redigo/redis"
	"github.com/stretchr/testify/assert"
)

var conn redis.Conn

func init() {
	var err error
	conn, err = redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatalln("Dial:", err)
	}
	log.Println("redis connect success!")
}

func TestInt(t *testing.T) {
	t.Cleanup(func() { conn.Close() })
	key := "keyInt"
	val := 12

	//set
	_, err := conn.Do("set", key, val)
	if err != nil {
		t.Fatal("Do set:", err)
	}

	//get
	reply, err := redis.Int(conn.Do("get", key))
	if err != nil {
		t.Fatal("Ints:", err)
	}

	assert.Equal(t, val, reply)
}

/*
ok
*/

func TestInts(t *testing.T) {
	key := "keyInts"
	val := []int{1, 2, 3}

	//set
	_, err := conn.Do("set", key, val)
	if err != nil {
		t.Fatal("Do set:", err)
	}

	//get
	reply, err := redis.Ints(conn.Do("get", key))
	if err != nil {
		t.Fatal("Ints:", err)
	}

	assert.Equal(t, val, reply)
}

/*
Ints: redigo: unexpected type for Ints, got type []uint8
*/

func TestIntMap(t *testing.T) {
	key := "keyIntMap"
	val := map[string]int{}
	val["a"] = 10000000
	val["b"] = 2

	//set
	_, err := conn.Do("set", key, val)
	if err != nil {
		t.Fatal("Do set:", err)
	}

	//get
	reply, err := redis.IntMap(conn.Do("get", key))
	if err != nil {
		t.Fatal("IntMap:", err)
	}

	assert.Equal(t, val, reply)
}

/*
IntMap: redigo: unexpected type for Values, got type []uint8
*/

func TestString(t *testing.T) {
	key := "keyString"
	val := "stringVal"

	//set
	_, err := conn.Do("set", key, val)
	if err != nil {
		t.Fatal("Do set:", err)
	}

	//get
	reply, err := redis.String(conn.Do("get", key))
	if err != nil {
		t.Fatal("String:", err)
	}

	assert.Equal(t, val, reply)
}

package redigo

import (
	"sort"
	"testing"

	"github.com/garyburd/redigo/redis"
	"github.com/stretchr/testify/assert"
)

func TestMSet(t *testing.T) {
	t.Cleanup(func() { conn.Close() })
	key1 := "hello"
	val1 := "HELLO"
	key2 := "world"
	val2 := "WORLD"
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

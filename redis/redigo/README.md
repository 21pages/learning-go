# API

## dial

```go
// Dial connects to the Redis server at the given network and
// address using the specified options.
func Dial(network, address string, options ...DialOption) (Conn, error)

// DialTimeout acts like Dial but takes timeouts for establishing the
// connection to the server, writing a command and reading a reply.
//
// Deprecated: Use Dial with options instead.
func DialTimeout(network, address string, connectTimeout, readTimeout, writeTimeout time.Duration) (Conn, error) {
	return Dial(network, address,
		DialConnectTimeout(connectTimeout),
		DialReadTimeout(readTimeout),
		DialWriteTimeout(writeTimeout))
}
```



## conn

```go
type Conn interface {
	// Close closes the connection.
	Close() error

	// Err returns a non-nil value when the connection is not usable.
	Err() error

	// Do sends a command to the server and returns the received reply.
	Do(commandName string, args ...interface{}) (reply interface{}, err error)

	// Send writes the command to the client's output buffer.
	Send(commandName string, args ...interface{}) error

	// Flush flushes the output buffer to the Redis server.
	Flush() error

	// Receive receives a single reply from the Redis server
	Receive() (reply interface{}, err error)
}
```



# 使用

## 连接

```go
//默认端口6379
conn, err := redis.Dial("tcp", ":6379")
if err != nil {
log.Fatalln("Dial:", err)
}
defer conn.Close()
log.Println("redis connect success")
```

# 问题

## []int -> []uint8
```go
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
```
output:
```
Ints: redigo: unexpected type for Ints, got type []uint8
```
int的slice, map存进去, 取出来就变uint8.  
用命令行读没问题, 可能是库的问题
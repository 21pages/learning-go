# 安装运行


## 安装
[https://github.com/tporadowski/redis/releases](https://github.com/tporadowski/redis/releases)

## 运行

服务器, 默认后台程序

```shell
redis-server.exe redis.window.conf
```

客户端

```shell
redis-cli.exe -h 127.0.0.1 -p 6379
127.0.0.1:6379>ping
PONG
```

- -a: password

- --raw:避免中文乱码

## 数据类型

key:只能是string, string是二进制.

val:所有的数据类型也都是基于string

如果要改数据类型, 需要先del key

| type                | feature                               | additional                                                   | apply                                                        | example                   |
| ------------------- | ------------------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------- |
| string字符串        | string:string键值对                   | 可以是普通字符串, 可以是数字, 也可以是任意二进制, 最大512M, 二进制安全 |                                                              | set abc 123               |
| hash字典            | 键值对集合, Map                       | 适合存储对象,并且可以像数据库中update一个属性一样只修改某一项属性值 | 存储、读取、修改用户属性                                     | hset abc hello 1 world  2 |
| list列表            | 双向链表                              | 可以从两头添加,列表最多可存储 2^32 - 1 元素                  | 1,最新消息排行等功能(比如朋友圈的时间线) 2,消息队列          | lpush abc 1 2 3           |
| set集合             | 哈希表, key:array, 无序存储的list     | 值不重复;内部由hash实现, 复杂度是 O(1);为集合提供了求交集、并集、差集等操作 | 1、共同好友 2、利用唯一性,统计访问网站的所有独立ip 3、好友推荐时,根据tag求交集,大于某个阈值就可以推荐 | sadd abc 1 2 3            |
| sorted set有序集合  | 有序哈希表, key:array, 有序存储的list | 每个value有个分数score, 按照score从小到大排序                | 1、排行榜 2、带权重的消息队列                                | zadd abc 1 1 2 2          |
| HyperLogLog基数统计 |                                       |                                                              | 基数:比如数据集 {1, 3, 5, 7, 5, 7, 8}， 那么这个数据集的基数集为 {1, 3, 5 ,7, 8}, 基数(不重复元素)为5 | pfcount abc               |
| stream流            | 持久化                                | 相对于pubsub,可以持久化,主备复制                             | 消息队列                                                     | xadd abc * a 1 b 2        |




# help命令

## help
```shell
127.0.0.1:6379> help
redis-cli 5.0.14 (git:a7c01ef4)
To get help about Redis commands type:
      "help @<group>" to get a list of commands in <group>
      "help <command>" for help on <command>
      "help <tab>" to get a list of possible help topics
      "quit" to exit
```
## help @group
```shell
help @string
help @hash
help @list
help @set
help @sorted_set
help @hyperloglog
help @stream
help @connection
help @server
```

## help  command

```shell
127.0.0.1:6379> help set

  SET key value [expiration EX seconds|PX milliseconds] [NX|XX]
  summary: Set the string value of a key
  since: 1.0.0
  group: string
```




# key

## 常用命令

参照[https://redis.io/commands](https://redis.io/commands)

| 命令                                    | 说明                                                       |
| --------------------------------------- | ---------------------------------------------------------- |
| del key                                 | 在 key 存在时删除 key                                      |
| dump key                                | 序列化给定 key ，并返回被序列化的值                        |
| exists key                              | 检查给定 key 是否存在                                      |
| expire key second                       | 为给定 key 设置过期时间，以秒计                            |
| expireat key timestamp                  | 设置key的过期时间为 UNIX 时间戳                            |
| pexpire key millisecond                 | 毫秒过期                                                   |
| pexpireat key timestamp-milli           | 毫秒时间戳过期                                             |
| keys pattern                            | 查找所有符合给定模式( pattern)的 key                       |
| move key db                             | 将当前数据库的 key 移动到给定的数据库 db 当中              |
| persist key                             | 移除 key 的过期时间，key 将持久保持                        |
| pttl key                                | 以毫秒为单位返回 key 的剩余的过期时间                      |
| ttl key                                 | 以秒为单位，返回给定 key 的剩余生存时间(TTL, time to live) |
| randomkey                               | 从当前数据库中随机返回一个 key                             |
| rename key newkey                       | 修改 key 的名称                                            |
| type key                                | 返回 key 所储存的值的类型                                  |
| renamenx key newkey                     | 仅当 newkey 不存在时，将 key 改名为 newkey                 |
| scan cursor [match pattern] [count cnt] | 迭代数据库中的数据库键                                     |

注意: redis是单线程的, keys命令会阻塞数据库, 可以使用scan, 只是会重复

## 示例

```shell
127.0.0.1:6379> keys *
1) "b"
2) "ab"
3) "c"
4) "oneSlice"
5) "a"
```



# 字符串string

## 常用命令

```shell
127.0.0.1:6379> help @string

  APPEND key value
  summary: Append a value to a key
  //如果 key 已经存在并且是一个字符串， APPEND 命令将指定的 value 追加到该 key 原来值（value）的末尾

  DECR key
  summary: Decrement the integer value of a key by one
  //数字值减1

  DECRBY key decrement
  summary: Decrement the integer value of a key by the given number
  //数字值减decrement

  GET key
  summary: Get the value of a key
  //读值

  GETSET key value
  summary: Set the string value of a key and return its old value
  //设置值并返回旧值

  INCR key
  summary: Increment the integer value of a key by one
  //数字值加1

  INCRBY key increment
  summary: Increment the integer value of a key by the given amount
  //数字值加increment

  INCRBYFLOAT key increment
  summary: Increment the float value of a key by the given amount
  //数字值加increment(float)

  MGET key [key ...]
  summary: Get the values of all the given keys
  //读取多个值

  MSET key value [key value ...]
  summary: Set multiple keys to multiple values
  //设置多个值

  MSETNX key value [key value ...]
  summary: Set multiple keys to multiple values, only if none of the keys exis
  //都不存在才设置

  PSETEX key milliseconds value
  summary: Set the value and expiration in milliseconds of a key
  //设置值,毫秒过期

  SET key value [expiration EX seconds|PX milliseconds] [NX|XX]
  summary: Set the string value of a key
  //设置值,并可同时设置过期时间, 判断是否存在才设置

  SETEX key seconds value
  summary: Set the value and expiration of a key
  //设置值,秒过期

  SETNX key value
  summary: Set the value of a key, only if the key does not exist
  //不存在时设置值

  STRLEN key
  summary: Get the length of the value stored in a key
  //值长度
```

## 示例

- set
```shell
set key value [expiration EX seconds|PX milliseconds] [NX|XX]
EX:过期秒; PX:过期毫秒
NX:set if not exist; XX: set only if exist

127.0.0.1:6379> set abc 123 EX 30 NX
OK
127.0.0.1:6379> set abc 123 EX 30 NX
(nil)
127.0.0.1:6379> set abc 1234 EX 30 XX
OK
127.0.0.1:6379> get abc
"1234"
127.0.0.1:6379> del abc
(integer) 0
127.0.0.1:6379> set abc 1234 EX 30 XX
(nil)
127.0.0.1:6379> get abc
```
# 字典hash

## 常用命令

```shell
127.0.0.1:6379> help @hash

  HDEL key field [field ...]
  summary: Delete one or more hash fields
  //删除一个或多个哈希表字段

  HEXISTS key field
  summary: Determine if a hash field exists
  //字段是否存在

  HGET key field
  summary: Get the value of a hash field
  //读字段值

  HGETALL key
  summary: Get all the fields and values in a hash
  //读key的所有字段, 键和值

  HINCRBY key field increment
  summary: Increment the integer value of a hash field by the given number
  //字段值加increment

  HINCRBYFLOAT key field increment
  summary: Increment the float value of a hash field by the given amount
  //字段值加increment(float)

  HKEYS key
  summary: Get all the fields in a hash
  //读key的所有字段的键

  HLEN key
  summary: Get the number of fields in a hash
  //读key的字段数量

  HMGET key field [field ...]
  summary: Get the values of all the given hash fields
  //读多个字段的值

  HMSET key field value [field value ...]
  summary: Set multiple hash fields to multiple values
  //设置多个字段的值

  HSCAN key cursor [MATCH pattern] [COUNT count]
  summary: Incrementally iterate hash fields and associated values
  //迭代读字段

  HSET key field value
  summary: Set the string value of a hash field
  //设置单个字段值

  HSETNX key field value
  summary: Set the value of a hash field, only if the field does not exist
  //当字段不存在时设置值

  HSTRLEN key field
  summary: Get the length of the value of a hash field
  //字段值的长度

  HVALS key
  summary: Get all the values in a hash
  //获取key所有的字段值
```

## 示例

```shell
127.0.0.1:6379> del abc
(integer) 1
127.0.0.1:6379> hmset abc a 1 b 2
OK
127.0.0.1:6379> hget abc a
"1"
127.0.0.1:6379> hgetall abc
1) "a"
2) "1"
3) "b"
4) "2"
127.0.0.1:6379> hkeys abc
1) "a"
2) "b"
127.0.0.1:6379> hset abc c 3
(integer) 1
127.0.0.1:6379> hdel abc b
(integer) 1
127.0.0.1:6379> hvals abc
1) "1"
2) "3"
127.0.0.1:6379> hscan abc 0
1) "0"
2) 1) "a"
   2) "1"
   3) "c"
```

# 列表list

## 常用命令

```shell
127.0.0.1:6379> help @list

  BLPOP key [key ...] timeout
  summary: Remove and get the first element in a list, or block until one is available
  //block left pop, or timeout

  BRPOP key [key ...] timeout
  summary: Remove and get the last element in a list, or block until one is available
  //block right pop, or timeout

  BRPOPLPUSH source destination timeout
  summary: Pop a value from a list, push it to another list and return it; or block until one is available
  //block right pop source, then push the pop one to desktination, or timeout

  LINDEX key index
  summary: Get an element from a list by its index
  //read from left by index

  LINSERT key BEFORE|AFTER pivot value
  summary: Insert an element before or after another element in a list
  //左边插入值

  LLEN key
  summary: Get the length of a list
  //list长度

  LPOP key
  summary: Remove and get the first element in a list
 //left pop

  LPUSH key value [value ...]
  summary: Prepend one or multiple values to a list
  //left push

  LPUSHX key value
  summary: Prepend a value to a list, only if the list exists
  //left push if key exists

  LRANGE key start stop
  summary: Get a range of elements from a list
  //left range

  LREM key count value
  summary: Remove elements from a list
  //left remove

  LSET key index value
  summary: Set the value of an element in a list by its index
  //left set

  LTRIM key start stop
  summary: Trim a list to the specified range
  //left trim

  RPOP key
  summary: Remove and get the last element in a list
  //right pop

  RPOPLPUSH source destination
  summary: Remove the last element in a list, prepend it to another list and return it
  // right pop source ,then left push to destination

  RPUSH key value [value ...]
  summary: Append one or multiple values to a list
  //right push

  RPUSHX key value
  summary: Append a value to a list, only if the list exists
  //right push if list exists
```

## 示例

```shell
127.0.0.1:6379> rpushx abc 1
(integer) 0
127.0.0.1:6379> lset abc
(error) ERR wrong number of arguments for 'lset' command
127.0.0.1:6379> rpush abc 1
(integer) 1
127.0.0.1:6379> lpush abc 0
(integer) 2
127.0.0.1:6379> lindex abc 1
"1"
127.0.0.1:6379> lpop abc
"0"
127.0.0.1:6379> lindex abc 0
"1"
```

# 集合set

## 常用命令

```shell
127.0.0.1:6379> help @set

  SADD key member [member ...]
  summary: Add one or more members to a set
  //添加

  SCARD key
  summary: Get the number of members in a set
  //数量

  SDIFF key [key ...]
  summary: Subtract multiple sets
  //差集

  SDIFFSTORE destination key [key ...]
  summary: Subtract multiple sets and store the resulting set in
  //差集存储

  SINTER key [key ...]
  summary: Intersect multiple sets
  //交集

  SINTERSTORE destination key [key ...]
  summary: Intersect multiple sets and store the resulting set in
  //交集存储

  SISMEMBER key member
  summary: Determine if a given value is a member of a set
  //containes

  SMEMBERS key
  summary: Get all the members in a set
  //所有成员

  SMOVE source destination member
  summary: Move a member from one set to another
  //移动

  SPOP key [count]
  summary: Remove and return one or multiple random members from
  //pop

  SRANDMEMBER key [count]
  summary: Get one or multiple random members from a set
  //随机读取

  SREM key member [member ...]
  summary: Remove one or more members from a set
  //remove

  SSCAN key cursor [MATCH pattern] [COUNT count]
  summary: Incrementally iterate Set elements
  //迭代

  SUNION key [key ...]
  summary: Add multiple sets
  //并集

  SUNIONSTORE destination key [key ...]
  summary: Add multiple sets and store the resulting set in a key
  //并集存储
```

## 示例

```shell
127.0.0.1:6379> sadd a 1 2 3
(integer) 3
127.0.0.1:6379> scard a
(integer) 3
127.0.0.1:6379> sadd b 3 4 5
(integer) 3
127.0.0.1:6379> sdiff a b
1) "1"
2) "2"
127.0.0.1:6379> sdiffstore c b a
(integer) 2
127.0.0.1:6379> smembers c
1) "4"
2) "5"
127.0.0.1:6379> sinter a b
1) "3"
127.0.0.1:6379> sunion a b
1) "1"
2) "2"
3) "3"
4) "4"
5) "5"
```

# 有序集合 sorted_set

## 常用命令

```shell
127.0.0.1:6379> help @sorted_set

  BZPOPMAX key [key ...] timeout
  summary: Remove and return the member with the highest score from one or more sorted sets, or block until one is available
  //block zset pop max scored

  BZPOPMIN key [key ...] timeout
  summary: Remove and return the member with the lowest score from one or more sorted sets, or block until one is available
  //block zset pop min scored

  ZADD key [NX|XX] [CH] [INCR] score member [score member ...]
  summary: Add one or more members to a sorted set, or update its score if it already exists
  //添加

  ZCARD key
  summary: Get the number of members in a sorted set
  //数量

  ZCOUNT key min max
  summary: Count the members in a sorted set with scores within the given values
  //score范围内数量

  ZINCRBY key increment member
  summary: Increment the score of a member in a sorted set
  //增加member的score

  ZINTERSTORE destination numkeys key [key ...] [WEIGHTS weight] [AGGREGATE SUM|MIN|MAX]
  summary: Intersect multiple sorted sets and store the resulting sorted set in a new key
  //交集存储

  ZLEXCOUNT key min max
  summary: Count the number of members in a sorted set between a given lexicographical range
  //词典序范围内数量

  ZPOPMAX key [count]
  summary: Remove and return members with the highest scores in a sorted set
  //remove and return max scored

  ZPOPMIN key [count]
  summary: Remove and return members with the lowest scores in a sorted set
  //remove and return min scored

  ZRANGE key start stop [WITHSCORES]
  summary: Return a range of members in a sorted set, by index
  //读取范围
  
  ZREM key member [member ...]
  summary: Remove one or more members from a sorted set
  //remove

  ZREMRANGEBYSCORE key min max
  summary: Remove all members in a sorted set within the given scores
  //移除score

  ZREVRANGE key start stop [WITHSCORES]
  summary: Return a range of members in a sorted set, by index, with scores ordered from high to low
  //移除score范围

  ZSCAN key cursor [MATCH pattern] [COUNT count]
  summary: Incrementally iterate sorted sets elements and associated scores
  //迭代

  ZSCORE key member
  summary: Get the score associated with the given member in a sorted set
  //读取score

  ZUNIONSTORE destination numkeys key [key ...] [WEIGHTS weight] [AGGREGATE SUM|MIN|MAX]
  summary: Add multiple sorted sets and store the resulting sorted set in a new key
  //并集存储
```

## 示例

```shell
127.0.0.1:6379> del abc
(integer) 0
127.0.0.1:6379> zadd abc 1 a 2 b 3 c
(integer) 3
127.0.0.1:6379> zcount abc 1 3
(integer) 3
127.0.0.1:6379> zpopmax abc 2
1) "c"
2) "3"
3) "b"
4) "2"
127.0.0.1:6379> zscan abc 0
1) "0"
2) 1) "a"
   2) "1"
127.0.0.1:6379>
```

# 基数统计hyperloglog

## 常用命令

```shell
127.0.0.1:6379> help @hyperloglog

  PFADD key element [element ...]
  summary: Adds the specified elements to the specified HyperLogLog.
  //添加元素

  PFCOUNT key [key ...]
  summary: Return the approximated cardinality of the set(s) observed by the HyperLogLog at key(s).
  //统计基数

  PFMERGE destkey sourcekey [sourcekey ...]
  summary: Merge N different HyperLogLogs into a single one.
  //合并
```



## 示例

```shell
127.0.0.1:6379> del abc
(integer) 1
127.0.0.1:6379> pfadd abc 1 2 2 3
(integer) 1
127.0.0.1:6379> pfcount abc
(integer) 3
127.0.0.1:6379> pfadd def 2
(integer) 1
127.0.0.1:6379> pfmerge x abc def
OK
127.0.0.1:6379> pfcount x
(integer) 3
```

# 发布订阅pubsub

## 常用命令

```shell
127.0.0.1:6379> help @pubsub

  PSUBSCRIBE pattern [pattern ...]
  summary: Listen for messages published to channels matching the given patterns
  //pattern subscribe, 规则订阅

  PUBLISH channel message
  summary: Post a message to a channel
  //发布

  PUBSUB subcommand [argument [argument ...]]
  summary: Inspect the state of the Pub/Sub subsystem
  //查看状态

  PUNSUBSCRIBE [pattern [pattern ...]]
  summary: Stop listening for messages posted to channels matching the given patterns
  //pattern unsubscribe, 规则取消订阅

  SUBSCRIBE channel [channel ...]
  summary: Listen for messages published to the given channels
  //订阅

  UNSUBSCRIBE [channel [channel ...]]
  summary: Stop listening for messages posted to the given channels
  //取消订阅
```

## 示例

client1:

```shell
127.0.0.1:6379> subscribe abc
Reading messages... (press Ctrl-C to quit)
1) "subscribe"
2) "abc"
3) (integer) 1
```

client2:

```shell
127.0.0.1:6379> publish abc "hello world"
(integer) 1
```

client1:

```shell
127.0.0.1:6379> subscribe abc
Reading messages... (press Ctrl-C to quit)
1) "subscribe"
2) "abc"
3) (integer) 1
1) "message"
2) "abc"
3) "hello world"
```

# 事务transactions

## 常用命令

```shell
127.0.0.1:6379> help @transactions

  DISCARD -
  summary: Discard all commands issued after MULTI
  //放弃执行事务命令

  EXEC -
  summary: Execute all commands issued after MULTI
  //开始执行事务命令

  MULTI -
  summary: Mark the start of a transaction block
  //事务块开始

  UNWATCH -
  summary: Forget about all watched keys
  //取消对所有key的监视

  WATCH key [key ...]
  summary: Watch the given keys to determine execution of the MULTI/EXEC block
  //监视key, 如果key有变化, 事务将被打断
```

## 示例

```shell
127.0.0.1:6379> multi
OK
127.0.0.1:6379> del abc
QUEUED
127.0.0.1:6379> set abc 123
QUEUED
127.0.0.1:6379> set def 456
QUEUED
127.0.0.1:6379> exec
1) (integer) 0
2) OK
3) OK
```

# stream

## 常用命令

```shell
127.0.0.1:6379> help @stream

  XACK key group ID [ID ...]
  summary: Marks a pending message as correctly processed, effectively removing it from the pending entries list of the consumer group. Return value of the command is the number of messages successfully acknowledged, that is, the IDs we were actually able to resolve in the PEL.
  //将消息标记为"已处理"

  XADD key ID field string [field string ...]
  summary: Appends a new entry to a stream
  //添加消息到末尾

  XCLAIM key group consumer min-idle-time ID [ID ...] [IDLE ms] [TIME ms-unix-time] [RETRYCOUNT count] [force] [justid]
  summary: Changes (or acquires) ownership of a message in a consumer group, as if the message was delivered to the specified consumer.
  //转移消息的归属权

  XDEL key ID [ID ...]
  summary: Removes the specified entries from the stream. Returns the number of items actually deleted, that may be different from the number of IDs passed in case certain IDs do not exist.
  //删除消息

  XGROUP [CREATE key groupname id-or-$] [SETID key id-or-$] [DESTROY key groupname] [DELCONSUMER key groupname consumername]
  summary: Create, destroy, and manage consumer groups.
  //产生,删除,管理消费组

  XINFO [CONSUMERS key groupname] [GROUPS key] [STREAM key] [HELP]
  summary: Get information on streams and consumer groups
  //查看流和消费者组的相关信息

  XLEN key
  summary: Return the number of entires in a stream
  // 获取流包含的元素数量

  XPENDING key group [start end count] [consumer]
  summary: Return information and entries from a stream consumer group pending entries list, that are messages fetched but never acknowledged.
  //显示待处理消息的相关信息

  XRANGE key start end [COUNT count]
  summary: Return a range of elements in a stream, with IDs matching the specified IDs interval
  //获取消息列表，会自动过滤已经删除的消息

  XREAD [COUNT count] [BLOCK milliseconds] STREAMS key [key ...] ID [ID ...]
  summary: Return never seen elements in multiple streams, with IDs greater than the ones reported by the caller for each stream. Can block.
  //以阻塞或非阻塞方式获取消息列表

  XREADGROUP GROUP group consumer [COUNT count] [BLOCK milliseconds] STREAMS key [key ...] ID [ID ...]
  summary: Return new entries from a stream using a consumer group, or access the history of the pending entries for a given consumer. Can block.
  // 读取消费者组中的消息

  XREVRANGE key end start [COUNT count]
  summary: Return a range of elements in a stream, with IDs matching the specified IDs interval, in reverse order (from greater to smaller IDs) compared to XRANGE
  //反向获取消息列表，ID 从大到小

  XTRIM key MAXLEN [~] count
  summary: Trims the stream to (approximately if '~' is passed) a certain size
  //对流进行修剪，限制长度
```

## 示例

```shell
127.0.0.1:6379> xadd abc * a 1 b 2
"1641989367732-0"
```

# 连接

## 常用命令

```shell
127.0.0.1:6379> help @connection

  AUTH password
  summary: Authenticate to the server
  //密码

  ECHO message
  summary: Echo the given string
  //打印

  PING [message]
  summary: Ping the server
  //连接测试

  QUIT -
  summary: Close the connection
  //退出连接

  SELECT index
  summary: Change the selected database for the current connection
  //改变数据库

  SWAPDB index index
  summary: Swaps two Redis databases
  //交换数据库
```

## 示例

```shell
127.0.0.1:6379> ping
PONG
127.0.0.1:6379> quit

D:\program\Redis>
```

# 服务器

## 常用命令

```shell
127.0.0.1:6379> help @server
  BGREWRITEAOF -
  summary: Asynchronously rewrite the append-only file
  //异步执行一个 AOF（AppendOnly File） 文件重写操作

  BGSAVE -
  summary: Asynchronously save the dataset to disk
  //在后台异步保存当前数据库的数据到磁盘

  CLIENT GETNAME -
  summary: Get the current connection name
  //获取连接的名称

  CLIENT ID -
  summary: Returns the client ID for the current connection
  //连接id

  CLIENT KILL [ip:port] [ID client-id] [TYPE normal|master|slave|pubsub] [ADDR ip:port] [SKIPME yes/no]
  summary: Kill the connection of a client
  //关闭客户端连接

  CLIENT LIST -
  summary: Get the list of client connections
  //获取连接到服务器的客户端连接列表

  CLIENT PAUSE timeout
  summary: Stop processing commands from clients for some time
  //终止运行来自客户端

  CLIENT REPLY ON|OFF|SKIP
  summary: Instruct the server whether to reply to commands
  //命令开关

  CLIENT SETNAME connection-name
  summary: Set the current connection name
  //setname

  CLIENT UNBLOCK client-id [TIMEOUT|ERROR]
  summary: Unblock a client blocked in a blocking command from a different connection
  //unblock

  COMMAND -
  summary: Get array of Redis command details
  //获取 Redis 命令详情数组

  COMMAND COUNT -
  summary: Get total number of Redis commands
  //获取 Redis 命令总数

  COMMAND GETKEYS -
  summary: Extract keys given a full Redis command
  //获取给定命令的所有键

  COMMAND INFO command-name [command-name ...]
  summary: Get array of specific Redis command details
  //获取指定 Redis 命令描述的数组

  CONFIG GET parameter
  summary: Get the value of a configuration parameter
  //获取指定配置参数的值

  CONFIG RESETSTAT -
  summary: Reset the stats returned by INFO
  //重置 INFO 命令中的某些统计数据

  CONFIG REWRITE -
  summary: Rewrite the configuration file with the in memory configuration
  //对启动 Redis 服务器时所指定的 redis.conf 配置文件进行改写

  CONFIG SET parameter value
  summary: Set a configuration parameter to the given value
  //修改 redis 配置参数，无需重启

  DBSIZE -
  summary: Return the number of keys in the selected database
  //返回当前数据库的 key 的数量

  DEBUG OBJECT key
  summary: Get debugging information about a key
  //获取 key 的调试信息

  DEBUG SEGFAULT -
  summary: Make the server crash
  //让 Redis 服务崩溃

  FLUSHALL [ASYNC]
  summary: Remove all keys from all databases
  //删除所有数据库的所有key

  FLUSHDB [ASYNC]
  summary: Remove all keys from the current database
  //删除当前数据库的所有key

  INFO [section]
  summary: Get information and statistics about the server
  //获取 Redis 服务器的各种信息和统计数值

  LASTSAVE -
  summary: Get the UNIX time stamp of the last successful save to disk
  //返回最近一次 Redis 成功将数据保存到磁盘上的时间，以 UNIX 时间戳格式表示

  MEMORY DOCTOR -
  summary: Outputs memory problems report
  //生成内存报告

  MEMORY HELP -
  summary: Show helpful text about the different subcommands
  //内存帮助

  MEMORY MALLOC-STATS -
  summary: Show allocator internal stats
  //内存分配信息

  MEMORY PURGE -
  summary: Ask the allocator to release memory
  //释放内存

  MEMORY STATS -
  summary: Show memory usage details
  //内存使用信息

  MEMORY USAGE key [SAMPLES count]
  summary: Estimate the memory usage of a key
  since: 4.0.0

  MONITOR -
  summary: Listen for all requests received by the server in real time
  //实时打印出 Redis 服务器接收到的命令，调试用

  REPLICAOF host port
  summary: Make the server a replica of another instance, or promote it as master.
  //将当前服务器转变为指定服务器的从属服务器(slave server)

  ROLE -
  summary: Return the role of the instance in the context of replication
  //返回主从实例所属的角色

  SAVE -
  summary: Synchronously save the dataset to disk
  //同步保存数据到硬盘

  SHUTDOWN [NOSAVE|SAVE]
  summary: Synchronously save the dataset to disk and then shut down the server
  //异步保存数据到硬盘，并关闭服务器

  SLAVEOF host port
  summary: Make the server a replica of another instance, or promote it as master. Deprecated starting with Redis 5. Use REPLICAOF instead.
  //将当前服务器转变为指定服务器的从属服务器(slave server)

  SLOWLOG subcommand [argument]
  summary: Manages the Redis slow queries log
  //管理 redis 的慢日志

  SYNC -
  summary: Internal command used for replication
  //用于复制功能(replication)的内部命令

  TIME -
  summary: Return the current server time
  //当前时间
```

## 示例

```shell
127.0.0.1:6379> save
OK
127.0.0.1:6379> set abc 123
OK
127.0.0.1:6379> flushdb
OK
127.0.0.1:6379> get abc
(nil)
```


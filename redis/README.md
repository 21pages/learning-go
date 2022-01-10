# redis介绍

## 特点
* 持久化:Redis支持数据的持久化，可以将内存中的数据保存在磁盘中，重启的时候可以再次加载进行使用
* 存储数据结构:Redis不仅仅支持简单的key-value类型的数据，同时还提供string、list（链表）、set（集合）、hash表等数据结构的存储。
* 数据备份: Redis支持数据的备份，即master-slave模式的数据备份

## 优势
相对于sql数据库
* 性能极高: 读写速度快, 适合做缓存
* 丰富的数据类型: Redis支持二进制案例的 Strings, Lists, Hashes, Sets 及 Ordered Sets 数据类型操作
* 原子 – Redis的所有操作都是原子性的，意思就是要么成功执行要么失败完全不执行。单个操作是原子性的。多个操作也支持事务，即原子性，通过MULTI和EXEC指令包起来
* 丰富的特性 – Redis还支持 publish/subscribe, 通知, key 过期等等特性

相对于其它key-value数据库
* Redis有着更为复杂的数据结构并且提供对他们的原子性操作
* 相比在磁盘上相同的复杂的数据结构，在内存中操作起来非常简单，这样Redis可以做很多内部复杂性很强的事情
* 磁盘格式方面他们是紧凑的以追加的方式产生的，因为他们并不需要进行随机访问。

## 安装

## redigo

```shell
go get -u -v github.com/garyburd/redigo/redis
```

## go-redis

```shell
go get -u -v github.com/go-redis/redis
```




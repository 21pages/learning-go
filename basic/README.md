# 值和引用

* 值:数字类型,结构体, 数组
* 引用:slice, map, channel, string, interface

# 底层实现
* slice: 数组, 数组指针+len+cap
* map: hash
* channel:循环链表,epoll

# range
* slice: k,v
* map :k, v
* channel: v, 不close会一直阻塞读取

# make和new
1. make可以传大小, 容量.new不能传参
2. make只用于slice、map,channel，而new用于各种类型
3. make返回引用,new返回指针。
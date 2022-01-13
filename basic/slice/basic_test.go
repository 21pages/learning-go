package basic_test

import (
	"fmt"
	"testing"
)

/*
子切片是父切片的引用

[102 203]
[0 1 102 203 4 5]
*/
func TestChild(t *testing.T) {
	data := [...]int{0, 1, 2, 3, 4, 5}

	s := data[2:4]
	s[0] += 100
	s[1] += 200

	t.Log(s)
	t.Log(data)
}

/*
&s地址不会变, 但是超过cap后, 底层数组会变

0xc0000040a8
0xc000018300
0xc0000040a8
0xc000018300
0xc0000040a8
0xc00000c320
*/
func TestAppend(t *testing.T) {
	s := make([]int, 1, 2)
	fmt.Printf("%p\n", &s)
	fmt.Printf("%p\n", &s[0])
	s = append(s, 0)
	fmt.Printf("%p\n", &s)
	fmt.Printf("%p\n", &s[0])
	s = append(s, 0)
	fmt.Printf("%p\n", &s)
	fmt.Printf("%p\n", &s[0])
}

/*
超过cap后二倍分配

cap: 1 -> 2
cap: 2 -> 4
cap: 4 -> 8
cap: 8 -> 16
cap: 16 -> 32
cap: 32 -> 64
*/
func TestCap(t *testing.T) {
	s := make([]int, 0, 1)
	c := cap(s)

	for i := 0; i < 50; i++ {
		s = append(s, i)
		if n := cap(s); n > c {
			fmt.Printf("cap: %d -> %d\n", c, n)
			c = n
		}
	}
}

/*
copy(s1,s2):拷贝长度是最小的len
两个 slice 可指向同一底层数组，允许元素区间重叠

array data :  [0 1 2 3 4 5 6 7 8 9]
slice s1 : [8 9]
slice s2 : [0 1 2 3 4]
copied slice s1 : [8 9]
copied slice s2 : [8 9 2 3 4]
last array data :  [8 9 2 3 4 5 6 7 8 9]
*/
func TestCopy(t *testing.T) {
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("array data : ", data)
	s1 := data[8:]
	s2 := data[:5]
	fmt.Printf("slice s1 : %v\n", s1)
	fmt.Printf("slice s2 : %v\n", s2)
	copy(s2, s1)
	fmt.Printf("copied slice s1 : %v\n", s1)
	fmt.Printf("copied slice s2 : %v\n", s2)
	fmt.Println("last array data : ", data)
}

/*
字符串底层是数组, 可以转切片
但是字符串本身是不可变的

hello
hello world
你好世界
*/
func TestString(t *testing.T) {
	str := "hello world"
	s1 := str[0:5]
	fmt.Println(s1)

	s2 := []byte(str)
	s2[0] = 'H'
	fmt.Println(str)

	strcn := "你好世界"
	s3 := []rune(strcn)
	s3[0] = '我'
	fmt.Println(strcn)
}

/*
起始:结束:cap

[0 1 2 3 4 5] 6 8
*/
func TestComma(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	d2 := slice[:6:8]
	fmt.Println(d2, len(d2), cap(d2))
}

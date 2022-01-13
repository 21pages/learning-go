package channel_test

import (
	"fmt"
	"testing"
	"time"
)

/*
不close channel, range会一直阻塞
close 会打出finish

0
1
2
3
4
*/
func TestRange(t *testing.T) {
	r := make(chan int, 2)
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			r <- i
		}
		close(r)
	}()

	for i := range r {
		fmt.Println(i)
	}
	fmt.Println("finish")
}

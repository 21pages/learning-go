package channel_test

import (
	"fmt"
	"strconv"
	"testing"
)

func makeCakeAndSend1(cs chan string, flavor string, count int) {
	for i := 1; i <= count; i++ {
		cakeName := flavor + " Cake " + strconv.Itoa(i)
		cs <- cakeName //send a strawberry cake
	}
	close(cs)
}

func receiveCakeAndPack1(strbry_cs chan string, choco_cs chan string) {
	strbry_closed, choco_closed := false, false

	for {
		//if both channels are closed then we can stop
		if strbry_closed && choco_closed {
			return
		}
		fmt.Println("Waiting for a new cake ...")
		select {
		case cakeName, strbry_ok := <-strbry_cs:
			if !strbry_ok {
				strbry_closed = true
				fmt.Println(" ... Strawberry channel closed!")
			} else {
				fmt.Println("Received from Strawberry channel.  Now packing", cakeName)
			}
		case cakeName, choco_ok := <-choco_cs:
			if !choco_ok {
				choco_closed = true
				fmt.Println(" ... Chocolate channel closed!")
			} else {
				fmt.Println("Received from Chocolate channel.  Now packing", cakeName)
			}
		}
	}
}

func TestSelect(t *testing.T) {
	strbry_cs := make(chan string)
	choco_cs := make(chan string)

	//two cake makers
	go makeCakeAndSend1(choco_cs, "Chocolate", 3)   //make 3 chocolate cakes and send
	go makeCakeAndSend1(strbry_cs, "Strawberry", 4) //make 3 strawberry cakes and send

	//one cake receiver and packer
	receiveCakeAndPack1(strbry_cs, choco_cs) //pack all cakes received on these cake channels
}

/*
for {
	select {
	case val, ok <- chan:
		...
	default:
		...
	}
}
当chan有数据或者被关闭时, 执行case, 被关闭时ok=false
否则执行default
当chan被关闭时, <-将不会阻塞, 一直执行

output:
Waiting for a new cake ...
Received from Strawberry channel.  Now packing Strawberry Cake 1
Waiting for a new cake ...
Received from Strawberry channel.  Now packing Strawberry Cake 2
Waiting for a new cake ...
Received from Strawberry channel.  Now packing Strawberry Cake 3
Waiting for a new cake ...
Received from Chocolate channel.  Now packing Chocolate Cake 1
Waiting for a new cake ...
Received from Chocolate channel.  Now packing Chocolate Cake 2
Waiting for a new cake ...
Received from Chocolate channel.  Now packing Chocolate Cake 3
Waiting for a new cake ...
Received from Strawberry channel.  Now packing Strawberry Cake 4
Waiting for a new cake ...
 ... Chocolate channel closed!
Waiting for a new cake ...
 ... Chocolate channel closed!
Waiting for a new cake ...
 ... Strawberry channel closed!
*/

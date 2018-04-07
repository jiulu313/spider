package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for {
		//1 第一种方法,用两个变量接收
		//n, ok := <-c
		//if !ok {
		//	break
		//}
		//fmt.Printf("worker %d received %c \n", id, n)

		//2 第二种方法，用range
		for n := range c {
			fmt.Printf("worker %d received %c \n",id,n)
		}

	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)

	return c;
}

func chanDemo() {
	var channels [10]chan<- int

	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
}

func channelClose() {
	c := make(chan int)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	//chanDemo()
	//bufferedChannel()

	channelClose()
}

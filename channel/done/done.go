package main

import (
	"fmt"
	"time"
)

func doWorker(id int, c chan int, done chan bool) {
	for {
		//2 第二种方法，用range
		for n := range c {
			fmt.Printf("worker %d received %c \n", id, n)
			done <- true
		}
	}
}

type worker struct {
	in   chan int
	done chan bool
}

func createWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}

	go doWorker(id, w.in, w.done)

	return w;
}

func chanDemo() {
	var workers [10]worker

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	for _,worker := range workers{
		<-worker.done
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	for _,worker := range workers{
		<-worker.done
	}


	time.Sleep(time.Millisecond)
}

func main() {
	chanDemo()

}

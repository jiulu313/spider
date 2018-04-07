package main

import (
	"fmt"
	"runtime"
)

func sayHello(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}

	c <- sum
}

func demo1() {
	c := make(chan int, 2)
	c <- 1
	c <- 2

	fmt.Println(<- c)
	fmt.Println(<- c)

}

func main() {
	//a := []int{1, 3, 5, 7, 9}
	//c := make(chan int)
	//
	//go sum(a, c)
	//fmt.Println("sum=", <-c)


	demo1()

}

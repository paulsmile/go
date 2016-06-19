package main

import (
	"fmt"
)

var c chan int = make(chan int)

func foo(i int) chan int {
	go func() {
		c <- i
	}()
	return c
}

func main() {
	for i := 0; i < 10; i++ {
		foo(i)
	}
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}

	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	for v := range ch {
		fmt.Println(v)
		if len(ch) <= 0 {
			break
		}
	}
}

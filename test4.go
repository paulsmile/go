package main

import (
	"fmt"
	"runtime"
)

var ch chan int = make(chan int)

func loop() {
	for i := 0; i < 10; i++ {
		runtime.Gosched()
		fmt.Println(i)
	}
	ch <- 0
}

func main() {
	runtime.GOMAXPROCS(2)
	go loop()
	go loop()
	go loop()
	for i := 0; i < 3; i++ {
		<-ch
	}
}

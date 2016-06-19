package main 

import "fmt"
import "time"

func producer(c chan<- int) {
	go func()  {
		for i:=0;i<10000;i++ {
			c <- i
		}
	}()
}

func consumer(c <-chan int) {
	go func() {
		for i:=0;i<10000;i++ {
			v := <-c
			fmt.Println("received:", v)
		}
	}()
}


func main() {
	c := make(chan int, 1)
	producer(c)
	consumer(c)
	time.Sleep(100 * time.Millisecond)
}
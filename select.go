package main 

import (
	"fmt"
	"time"
)

func main() {
	c1:=make(chan string)
	c2:=make(chan string)

	go func() {
		c1 <- "receive hello world"
		
	}()

	go func() {
		c2 <- "recieve c2"
		time.Sleep(1*time.Second)
	}()

	for i:=0;i<2;i++ {
		select {
			case t1:=<-c1: fmt.Println(t1)
			case t2:=<-c2: fmt.Println(t2)
		}
	}
}
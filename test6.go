package main

import "fmt"
import "runtime"

func say(name string) {
	fmt.Println(name)
}

func main() {
	runtime.GOMAXPROCS(4)
	go say("try")
	go say("hello")
	for {

	}
}

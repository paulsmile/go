package main

import "fmt"

import "runtime"

func say(name string) {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func main() {
	runtime.GOMAXPROCS(4)
	go say("world")
	for {

	}

}

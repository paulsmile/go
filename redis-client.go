package main

import (
	"fmt"
	"github.com/hoisie/redis"
	"sync"
)

var (
	client *redis.Client
	mutex  sync.Mutex
)

func init() {
	mutex.Lock()
	defer mutex.Unlock()
	if client != nil {
		return
	}
	client = &redis.Client{
		Addr:        "localhost:6379",
		Db:          0,
		Password:    "",
		MaxPoolSize: 1000,
	}
}

func Sadd(key string, value []byte) (bool, error) {
	return client.Sadd(key, value)
}

func Score(key string) (int, error) {
	return client.Scard(key)
}

func main() {
	info, _ := Sadd("test", []byte("hello"))
	fmt.Println(info)
	score, err := Score("test")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("score:%d\n", score)
	}
	// for i := 0; i < 100000; i++ {
	// 	Sadd("test", []byte("test:"+string(i)))
	// }
	result, _ := Score("test")
	fmt.Println(result)

}

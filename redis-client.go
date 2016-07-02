package main

import (
	"fmt"
	"github.com/hoisie/redis"
	"sync"
	"time"
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

func Members(key string) ([][]byte, error) {
	return client.Smembers(key)
}

/**
 * @param key string
 * @param val []byte
 * @param t int64 means set the timeout
 * return error
 **/
func Set(key string, val []byte, t int64) error {
	return client.Setex(key, t, val)
}

func Get(key string) ([]byte, error) {
	return client.Get(key)
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

	members, _ := Members("test")
	fmt.Println(len(members))

	Set("aa", []byte("show me the code, jackson!!"), 3)

	value, _ := Get("aa")
	fmt.Println(string(value))

	// sleep 3 second later
	time.Sleep(3 * time.Second)

	aa_result, err := Get("aa")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(aa_result)
	}

}

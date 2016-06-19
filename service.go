package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	addr := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go HandleClient(conn)
	}

}

func HandleClient(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(1 * time.Minute))
	request := make([]byte, 128)
	defer conn.Close()
	for {
		read_len, _ := conn.Read(request)
		fmt.Println(read_len)
		fmt.Println(string(request[:read_len]))
		if read_len == 0 {
			break
		} else {
			conn.Write([]byte("see you again!"))
		}
		request = make([]byte, 128)

	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error:%s", err.Error())
		os.Exit(1)
	}
}

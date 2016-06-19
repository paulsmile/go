package main

import (
	"flag"
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls", "-al")
	ouput, err := cmd.Output()
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	fmt.Println(string(ouput))
	test := flag.String("test", "site.gizwits.com", "host")
	flag.Parse()
	fmt.Println(*test)
}

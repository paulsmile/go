package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	Name  string  `json:"item"`
	Other float32 `json:"amount"`
}

func main() {
	var detail Data
	detail.Name = "Jack"
	detail.Other = 32.12

	resp, err := json.Marshal(detail)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(resp))
}

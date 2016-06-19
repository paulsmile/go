package main 

import (
	"fmt"
)

type People interface {
	say(language string) string
}

type Person struct {
	name string
}

 func(p *Person) say(language string) string {
	return p.name + ", speak " + language
}

func Show(a int)  func(b int) int {
	return func(b int) int {
		return  a + b;
	}
}

func main() {
	
	var p1 People = new(Person)
	fmt.Println(p1)
	p1 = &Person {
		name: "Tom"}
	fmt.Println(p1)
	info := p1.say("Germany")

	var p2 People = new(Person)
	p2 = &Person {
		name: "Jack"}
	s := Show(1)
	for i:=0;i<10;i++ {
		fmt.Println(s(i))
	}
	t := Show(2)
	for i:=0;i<10;i+=2 {
		fmt.Println(t(i))
	}
	fmt.Println(info)
	fmt.Println(p2.say("English"))
}
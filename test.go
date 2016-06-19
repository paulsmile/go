package main

import "fmt"
import "math"

func NextNum() func() int {
	var i, j = 0, 1
	return func() int {
		tmp := i + j
		i, j = tmp, i
		return tmp
	}
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

type Cycle struct {
	Radius float64
}

type Result struct {
	Width, Height float64
}

func (c *Cycle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c *Cycle) Parimeter() float64 {
	return c.Radius * c.Radius
}

func (r *Result) Area() float64 {
	return 2 * (r.Width + r.Height)
}

func (r *Result) Parimeter() float64 {
	return r.Width * r.Height
}

type shape interface {
	Area() float64
	Parimeter() float64
}

func main() {
	var a = NextNum()
	for i := 0; i < 10; i++ {
		fmt.Println(a())
	}

	var g = fact(10)
	fmt.Println(g)
	c := Cycle{Radius: 4.2}
	r := Result{Width: 5.0, Height: 3.2}
	s1 := []shape{&c, &r}
	for s, _ := range s1 {
		fmt.Println(s)
	}
}

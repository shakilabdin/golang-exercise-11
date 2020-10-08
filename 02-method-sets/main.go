package main

import (
	"fmt"
)

func main() {

	p1 := person{
		"Shakil",
		"Abdin",
		25,
	}

	saySomething(&p1)
}

type person struct {
	first string
	last  string
	age   int
}

type human interface {
	hello()
}

func (p *person) hello() {
	fmt.Println("Hello, my name is", p.first, p.last)
}

func saySomething(h human) {
	h.hello()
}
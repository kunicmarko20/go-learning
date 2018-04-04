package main

import (
	"fmt"
)

type Person struct {
	Name string
	Address
}

type Person2 struct {
	Name string
	Address
}

type Address struct {
	Street string
	Number string
	City string
}

func (a Address) getAddress() string {
	return a.City + ", " + a.Street + " " + a.Number
}

func main() {
	p1 := Person{"name", Address{"street", "number", "city"}}
	var p2 Person2
	fmt.Println(p1.getAddress())
	fmt.Println(p2.getAddress())
}

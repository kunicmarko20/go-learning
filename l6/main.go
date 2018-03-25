package main

import "fmt"

func main() {
	fmt.Println(calc(11, 54, 23, 33, 88, 78))
}

func calc(num ...int) int{
	var highest int

	for _, v := range num {
		if v > highest {
			highest = v
		}
	}

	return highest
}
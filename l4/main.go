package main

import "fmt"

func main() {
	var a, b = greet("test", "test2")

	fmt.Println(a, b)
}

func greet(fname, lname string) (string, string) {
	return fname, lname
}

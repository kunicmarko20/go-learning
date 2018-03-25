package main

import "fmt"

func main() {
	// string type is set for both even if you set it to 2nd param
	greet(1, "test")
}

func greet(fname, lname string) {
	fmt.Println(fname, lname)
}

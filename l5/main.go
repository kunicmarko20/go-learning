package main

import "fmt"

func main() {
	fmt.Println(calc(1))
}

func calc(num int) (int, bool){
	return num / 2, num % 2 == 0
}
package main

import "fmt"

func main() {
	var slice = make([]int, 0, 50)

	for i := 0; i < 100; i++ {
		slice = append(slice, i)
		fmt.Println("L:", len(slice), "C:", cap(slice), "V:", slice[i])
	}
}

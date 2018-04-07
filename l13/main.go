package main

import (
	"fmt"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

func main() {
	waitGroup.Add(2)
	go foo()
	go bar()
	waitGroup.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("Foo:", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}

	waitGroup.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("Bar:", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}

	waitGroup.Done()
}

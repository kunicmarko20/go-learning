package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
	"sync/atomic"
)

var waitGroup sync.WaitGroup
var counter int64

func main() {
	waitGroup.Add(2)
	go incrementor("foo:")
	go incrementor("bar:")
	waitGroup.Wait()
}

func incrementor(word string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)

		atomic.AddInt64(&counter, 1);
		fmt.Println(word, i, "Counter:", counter)
	}

	waitGroup.Done()
}

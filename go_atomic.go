package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var a int64 = 0
var wg sync.WaitGroup

func f() {
	for i := 0; i < 10_000; i++ {
		atomic.AddInt64(&a, 1)
	}
	wg.Done()
}
func main() {
	const size = 100
	for i := 0; i < size; i++ {
		wg.Add(1)
		go f()
	}
	wg.Wait()
	fmt.Println(a)
}

package main

import (
	"fmt"
	"os"
	"time"
)

var exit bool = false

func main() {
	tick := time.Tick(1 * time.Second)
	abort := make(chan struct{})
	go func() {
		for {
			select {
			case <-tick:
				fmt.Println("Tick")
			case <-abort:
				fmt.Println("Exit")
				return
			}
		}
	}()
	os.Stdin.Read(make([]byte, 1))
	abort <- struct{}{}
}

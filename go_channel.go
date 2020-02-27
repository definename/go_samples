package main

import (
	"fmt"
)

func f(c chan int) {
	a := 0
	for i := 0; i < 10_000; i++ {
		a++
	}
	c <- a
}
func main() {
	// Sync with channel
	c := make(chan int)
	const size = 100
	for i := 0; i < size; i++ {
		go f(c)
	}
	res := 0
	for i := 0; i < size; i++ {
		res += <-c
	}
	fmt.Println(res)

	// Close channel
	numbers := make(chan int)
	doubles := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			numbers <- i
		}
		close(numbers)
	}()
	go func() {
		for v := range numbers {
			doubles <- 2 * v
		}
		close(doubles)
	}()
	for v := range doubles {
		fmt.Print(v, " ")
	}
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func consumer(i int, sem chan int) {
	for {
		c := <-sem
		fmt.Printf("\t\tCon:%d: %v\n", i, c)
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	}
}

func producer(i int, sem chan int) {
	for {
		c := rand.Intn(10)
		fmt.Printf("Pro:%d: %v\t\n", i, c)
		sem <- c
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	}
}

func main() {
	sem := make(chan int, 10)
	for i := 0; i < 2; i++ {
		go producer(i, sem)
		go consumer(i, sem)
	}
	for {
	}
}

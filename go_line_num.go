package main

import (
	"log"
)

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			log.Println("FizzBuzz")
		} else if i%3 == 0 {
			log.Println("Fizz")
		} else if i%5 == 0 {
			log.Println("Buzz")
		} else {
			log.Println(i)
		}
	}
}

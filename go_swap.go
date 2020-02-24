package main

import (
	"log"
)
func main() {
	arr := [5]int{1, 0, 2, 99, 3}
	min, minI := 0, 0
	max, maxI := 0, 0
	for i, v := range arr {
		if v < min {
			minI = i
		}
		if v > max {
			maxI = i
		}
		log.Println(i, v)
	}
	log.Println("Index:", minI, maxI)
}
package main

import (
	"log"
)

func swap(arr []int, index1 int, index2 int) {
	tmp := arr[index1]
	arr[index1] = arr[index2]
	arr[index2] = tmp
}

func main() {
	arr := [5]int{1, 0, 2, 99, 3}
	log.Println("arr", arr)
	min, minI := arr[0], 0
	max, maxI := arr[0], 0
	log.Println("init max:", max, "init min:", min)
	for i, v := range arr {
		if min > v {
			log.Println("min")
			minI = i
			min = v
		}
		if max < v {
			log.Println("max")
			maxI = i
			max = v
		}
		log.Println(i, v)
	}
	log.Println("min:", min, "index:", minI)
	log.Println("max:", max, "index:", maxI)

	swap(arr[:], minI, maxI)
	log.Println("arr", arr)
}

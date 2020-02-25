package main

import ("log")

func main() {
	arrA := [5]int{0, 1, 2, 4, 100}
	arrB := [4]int{5, 6, 7, 8}

	arrC := make([]int, len(arrA), len(arrA) + len(arrB))
	copy(arrC, arrA[:])
	log.Println(arrC)
	arrC = append(arrC[:1], 99)
	log.Println(arrC)

	// for i := 0; i < len(arrB); i++ {
	// 	toInsert := arrB[i]
	// 	for j := 0; j < len(arrC); j++ {
	// 		if toInsert > arrC[j] {
	// 			arrC = append(arrC[:j], toInsert)
	// 			break
	// 		}
	// 	}
	// }
	// log.Println(arrC)
}
package main

import "log"

func main() {
	arrA := [5]int{0, 1, 2, 4, 100}
	log.Println("arrA", arrA)
	arrB := [6]int{5, 6, 7, 8, -1, 8}
	log.Println("arrB", arrB)

	arrC := make([]int, len(arrA))
	copy(arrC, arrA[:])
	log.Println("arrC", arrC)
	for i := 0; i < len(arrB); i++ {
		toInsert := arrB[i]
		for j := 0; j < len(arrC); j++ {
			value := arrC[j]
			if toInsert == value {
				log.Println("Skip:", toInsert)
				break
			}
			if toInsert < value {
				log.Println("toInsert:", toInsert)
				arrC = append(arrC, 0)
				copy(arrC[j+1:], arrC[j:])
				arrC[j] = toInsert
				break
			}
		}
	}
	log.Println("merged:", arrC, "len:", len(arrC))
}

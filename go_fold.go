package main

import "log"

func fold(op func(int, int)int, s []int) int{
	res := 0
	for i := 0; i < len(s); i++ {
		res = op(res, s[i])
	}
	return res
}

func add(x int, y int) int {
	return x + y
}

func main() {
	s := []int{1, 2, 3}
	log.Println(fold(add, s))
}
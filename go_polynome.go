package main

import "log"

func getPolynomeFunc(s []int) func(int) int {
	return func(x int) int {
		res := 0
		for _, v := range s {
			res = res*x + v
		}
		return res
	}
}

func main() {
	s := []int{1, 2, 3}
	polynomeFunc := getPolynomeFunc(s)
	r := polynomeFunc(1)
	log.Println(r)
}

package main

import (
	"log"
)

// Function
func fact(n int) int {
	if n < 2 {
		return 1
	}
	return n * fact(n-1)
}

func divmod(x, y int) (int, int) {
	return x / y, x % y
}

func add(x, y int) (z int) {
	z = x + y
	return
}

func getIntPtr() *int {
	a := 1
	return &a
}

func sum(vs ...int) int {
	s := 0
	for _, v := range vs {
		s += v
	}
	return s
}

var z = func(x, y int) int {
	return x + y
}(2, 3)

type meter int
type kilometer int

func mToKm(v meter) kilometer {
	return kilometer(v / 1000)
}

func main() {
	log.Println(fact(5))

	d, r := divmod(5, 2)
	log.Println(d, r)

	log.Println("Add", add(5, 2))

	// Return ptr on local variable from function
	p := getIntPtr()
	*p += 3
	log.Println(*p)

	// Function pointer
	var fPtr func(int) int
	fPtr = fact
	log.Println("Fact:", fPtr(3))

	// Closure
	l := []func(){}
	for i := 0; i < 3; i++ {
		x := i
		l = append(l, func() { log.Println("closure", x) })
	}

	for _, f := range l {
		f()
	}

	// Unpack args
	log.Println("Unpack:", sum(1, 2, 3))
	sliceToUnpack := []int{1, 2, 3}
	log.Println("Unpack:", sum(sliceToUnpack...))

	// Default value
	log.Println("Def value:", z)

	// Defer
	defer log.Println("world1")
	defer log.Println("world2")
	log.Println("Hello")

	// Alias
	var m meter = 1000
	var k kilometer = mToKm(m)
	log.Println("meter to kilometer:", k)
}

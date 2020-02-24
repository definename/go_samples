package main

import (
	"fmt"
	"log"
	"unsafe"
)

func main() {
	// Constant and types.
	const (
		Flag0 int = 1 << iota
		Flag1
		Flag2 uint8 = 1
		Flag3
		a   int8   = 0b00000011
		b   int8   = 0b00000101
		c   bool   = true
		str string = "Go!"
	)

	fmt.Println(Flag0, Flag1, Flag2, Flag3)
	fmt.Printf("a:%v, type is:%[1]T, length is:%d\n", Flag2, unsafe.Sizeof(Flag2))
	fmt.Printf("%08b\n", a&^b)
	fmt.Printf("%t\n", c)
	fmt.Printf("%s %d\n", str, len(str))

	// Variables
	var i int = 99
	fmt.Println(i + 1)
	aa, bb := 1, 2
	bb, aa = aa, bb
	fmt.Println(aa, bb)

	if x := aa + 1; x > bb {
		fmt.Println("a was chosen", aa)
	} else {
		fmt.Println("b was chosen", bb)
	}

	// Loop
	for i := 32; i < 33; i++ {
		fmt.Printf("->%d\n", i)
	}

	arr := []int{11, 22}
	for index, value := range arr {
		log.Println("index:", index, "value", value)
	}

	sw := 55
	switch sw {
	case 55, 44, 33:
		fmt.Printf("=>%d\n", sw)
	case 50:
		fmt.Printf("=>%d\n", sw)
	}

	switch {
	case sw > 40:
		fmt.Printf(">40\n")
		fallthrough
	case sw > 100:
		fmt.Printf(">50\n")
	}
}

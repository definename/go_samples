package main

import (
	"fmt"
	"log"
)

func main() {
	// goto
	i := 0
makr1:
	i++
	if i < 5 {
		log.Println("mark1:", i)
		goto makr1
	}

	// break/continue to
mark2:
	for j := 0; j < 10; j++ {
		if j >= 5 {
			break mark2
		}
		log.Println("makr2:", j)
	}

	// array
	a := [4]int{0, 1, 2, 3: 99}
	a[0] = 100
	log.Println(a)
	b := [...]int{11, 22, 33, 44}
	log.Println(b)

	// slice
	sl := []int{0, 1, 2, 3, 4}
	log.Println("slice:", sl)
	slice1 := make([]int, 5, 7)
	copy(slice1, sl)
	log.Println("slice1:", slice1, len(slice1), cap(slice1))
	slice2 := slice1[2:4]
	slice2[0] = 99
	slice2 = append(slice2, 999)
	log.Println("slice2", slice2, "slice1", slice1)

	// strings
	str := "Привіт!"
	for _, v := range str {
		log.Printf("%v %[1]c", v)
	}

	multilineStr := `Привіт
світ`
	log.Println(multilineStr)

	test := 1
	fmtStr := fmt.Sprintf("%v+%[1]v", test)
	log.Println(fmtStr)

	// map
	m := make(map[string]int)
	m["a"] = 1
	log.Println(m, m["a"])
	v, ok := m["b"]

	log.Println(v, ok)
	if v, ok := m["a"]; ok {
		log.Println("Ok", v)
	}
	delete(m, "a")
	log.Println("map len is:", len(m))

	mm := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	for key, val := range mm {
		log.Println("key:", key, "value:", val)
	}

	// pointer
	data := 1
	ptr := &data
	log.Printf("%p %[1]T, %v\n", ptr, *ptr)
	var ptr1 *int = ptr
	*ptr1 = 2
	log.Println("data:", data)
	var ptr2 *int = new(int)
	*ptr2 = 3
	log.Println("ptr2:", *ptr2)
}

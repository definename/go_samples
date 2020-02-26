package main

import (
	"fmt"
	"log"
)

type BaseStar struct {
	name string
}

type Star struct {
	BaseStar
	distance      int
	constellation string
}

type Address struct {
	city   string
	street string
	house  int
}

func MakeAddress(city, street string, house int) Address {
	return Address{city, street, house}
}

func NewAddress(city, street string, house int) *Address {
	return &Address{city, street, house}
}

func (a Address) getCityV() string {
	return a.city
}
func (a *Address) getCityP() string {
	return a.city
}
func (a Address) setCityV(city string) {
	a.city = city
}
func (a *Address) setCityP(city string) {
	a.city = city
}

type Matrix struct {
	a [][]int
}

func NewMatrix(l int) *Matrix {
	m := &Matrix{}
	for i := 0; i < l; i++ {
		m.a = append(m.a, make([]int, l))
	}
	return m
}

func (m *Matrix) Fill(v int) {
	for i, row := range m.a {
		for j, _ := range row {
			m.a[i][j] = v
		}
	}
}

func (m *Matrix) MyString() string {
	return fmt.Sprintf("Matrix([%v ... %v])", m.a[0][0], m.a[len(m.a)-1][len(m.a)-1])
}

type Shape interface {
	getArea() float64
}

type Square struct {
	a float64
}

func (s Square) getArea() float64 {
	return s.a * s.a
}

type Circle struct {
	r float64
}

func (c Circle) getArea() float64 {
	return 3.14 * (c.r * c.r)
}

func printType(i interface{}) {
	fmt.Printf("%v, %[1]T\n", i)
}

type Some interface {
	do() int
}

type X int

func (x X) do() int {
	return int(x)
}

func execute(s Some) {
	fmt.Println("do:", s.do())
}

func execute_with_cast(s interface{}) {
	v, ok := s.(X)
	if !ok {
		fmt.Println("Incorrect type assertion")
	} else {
		fmt.Println("do with cast:", v.do())
	}
}

func go_panic() {
	defer func() {
		fmt.Println("Defer")
		recover()
	}()
	panic("Panic!")
}

func main() {
	// Struct
	aldebaran1 := Star{BaseStar{"Aldebaran"}, 65, "Taurus"}
	log.Println(aldebaran1)
	aldebaran2 := Star{
		BaseStar:      BaseStar{name: "Aldebaran"},
		distance:      65,
		constellation: "Taurus"}
	log.Println(aldebaran2)
	log.Println(aldebaran1 == aldebaran2)
	aldebaran3 := &Star{BaseStar{"Aldebaran"}, 65, "Taurus"}
	log.Println(aldebaran3)

	anonymStruct := struct {
		firstName string
		lastName  string
		salary    int
	}{
		firstName: "Harry",
		lastName:  "Potter",
		salary:    9999,
	}
	log.Println(anonymStruct)

	addr1 := MakeAddress("London", "Private drive", 4)
	addr2 := NewAddress("London", "Private drive", 4)
	log.Printf("%v %[1]T", addr1, addr1)
	log.Printf("%v %[1]T", addr2, addr2)

	// Value receiver
	log.Println(addr1.getCityV())
	log.Println(addr2.getCityP())

	addr1.setCityV("Kiev")
	addr2.setCityP("Kiev")
	log.Println(addr1.getCityV())
	log.Println(addr2.getCityP())

	m := NewMatrix(4)
	log.Println(m)
	m.Fill(1)
	log.Println(m)
	log.Println(m.MyString())

	// Interface
	var s Shape
	s = Square{5}
	fmt.Println(s.getArea())
	s = Circle{5}
	fmt.Println(s.getArea())

	printType(s)

	x := X(5)
	execute(x)
	execute(&x)
	execute_with_cast(x)

	go_panic()
	fmt.Println("No panic!")
}

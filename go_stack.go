package main

import (
	"fmt"
	"strings"
)

type Node struct {
	val  int
	next *Node
}

type Stack struct {
	top *Node
}

func (s *Stack) String() string {
	res := make([]string, 0)
	for t := s.top; t != nil; t = t.next {
		res = append(res, fmt.Sprintf("%d", t.val))
	}
	return strings.Join(res, "->")
}

func (s *Stack) Push(v int) {
	n := &Node{v, s.top}
	s.top = n
}

func (s *Stack) Pop() (int, bool) {
	if s.top != nil {
		n := s.top
		s.top = n.next
		return n.val, false
	}
	return 0, true
}

func (s *Stack) Get() (int, bool) {
	if s.top != nil {
		n := s.top
		return n.val, false
	}
	return 0, true
}

func main() {
	s := &Stack{nil}
	s.Push(10)
	s.Push(1)
	fmt.Println(s)

	v, err := s.Get()
	if !err {
		fmt.Printf("Got:%d\n", v)
	}

	for {
		v, err := s.Pop()
		if err {
			break
		}
		fmt.Printf("Pop:%d\n", v)
	}
}

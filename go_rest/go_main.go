package main

import (
	"fmt"
	"net/http"

	"go_rest/model"
)

type MyServer struct{}

func (s *MyServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		fmt.Fprintf(w, "message: Post")
	case "GET":
		fmt.Fprintf(w, "message: Get")
		model.GetAll()
	default:
		fmt.Fprintf(w, "message: Donkey")
	}
}

var (
	a int
	b int
)

func One() {
	i := 1
	j := 2
	a, b = i, j
	fmt.Println(a, b, i, j)
}

func main() {
	// One()
	// fmt.Println(a, b)

	model.OpenDb()

	http.Handle("/hello", &MyServer{})
	http.ListenAndServe(":8000", nil)

	model.CloseDb()
}

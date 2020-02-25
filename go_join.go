package main

import (
	"log"
	"strings"
)

func variadicJoin(sep string, params ...string) string {
	return strings.Join(params, sep)
}

func main() {
	origin := "to"
	log.Println(variadicJoin(origin, "1", "2", "3"))
}

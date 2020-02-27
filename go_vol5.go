package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

type Pair struct {
	Word string
	Count int
}
type PairList []Pair

// func (p PairList) Len() int { return len(p) }
// func (p PairList) Less(i, j int) bool { return p[i].Count < p[j].Count}
// func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i]}

func main() {
	data, err := ioutil.ReadFile("pg1661.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := strings.ToLower(string(data))
	reg, err := regexp.Compile("[0-9a-z]+")
	if err != nil {
		log.Fatal(err)
	}
	words := reg.FindAllString(text, -1)
	m := make(map[string]int)
	for _, word := range words {
		// if count, ok := m[word]; ok {
		// 	m[word] = count + 1
		// } else {
		// 	m[word] = 1
		// }
		m[word]++
	}
	// fmt.Println(m)

	counter := make(PairList, len(m))
	i := 0
	for word, count := range m {
		counter[i] = Pair{word, count}
		i++
	}
	// sort.Sort(counter)
	sort.Slice(counter, func (i, j int) bool {
		return counter[i].Count < counter[j].Count
	})
	fmt.Println(counter)
	f, err := os.Create("stat.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	for _, pair := range counter {
		f.WriteString(fmt.Sprintf("%v\t%v\n", pair.Word, pair.Count))
	}
}
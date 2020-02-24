package main

import (
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	// Date parsing
	layout := "2006.01.02"
	input := "2020.02.24"
	parsedTime, e := time.Parse(layout, input)
	if e == nil {
		log.Printf("Datetime parsed:%s\n", parsedTime)
	} else {
		log.Println("Failed to parse given time string")
	}

	// Week day to string
	if len(os.Args) < 2 {
		log.Fatal("Invalid arg was given")
	}

	weekDayNum, e := strconv.ParseInt(os.Args[1], 10, 8)
	if e == nil && weekDayNum >= 0 && weekDayNum < 7 {
		log.Println(time.Weekday(weekDayNum))
	} else {
		log.Fatal("Invalid week day format or out of range")
	}
}

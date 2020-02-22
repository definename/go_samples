package main

import (
	"log"
	"math"
	"os"
	"strconv"
)

func isValidHeronian(evenSide float64, oddSide float64) bool {
	log.Printf("Even:%f odd:%f", evenSide, oddSide)
	return evenSide == 1
}

func isTriangel(a float64, b float64, c float64) bool {
	valid := false
	valid = a < (b+c) && a > (b-c)
	if valid {
		valid = b < (a+c) && b > (a-c)
		if valid {
			valid = c < (a+b) && c > (a-b)
		}
	}
	return valid
}

func main() {
	evenSide := 0.0
	oddSide := 0.0
	if len(os.Args) < 3 {
		log.Fatal("Please provide a,b,c triangel sides.")
	}

	a, errA := strconv.ParseFloat(os.Args[1], 64)
	if errA != nil {
		log.Fatal("A is incorrect")
	}
	if math.Mod(a, 2) == 0 {
		evenSide++
	} else {
		oddSide++
	}

	b, errB := strconv.ParseFloat(os.Args[2], 64)
	if errB != nil {
		log.Fatal("B is incorrect")
	}
	if math.Mod(b, 2) == 0 {
		evenSide++
	} else {
		oddSide++
	}

	c, errC := strconv.ParseFloat(os.Args[3], 64)
	if errC != nil {
		log.Fatal("C is incorrect")
	}
	if math.Mod(c, 2) == 0 {
		evenSide++
	} else {
		oddSide++
	}

	log.Printf("Side a:%f b:%f c:%f", a, b, c)
	if isTriangel(a, b, c) {
		if isValidHeronian(evenSide, oddSide) {
			var semiPerimeter float64
			semiPerimeter = (a + b + c) / 2
			log.Println("Heronian semiperimeter is:", semiPerimeter)

			area := math.Sqrt(semiPerimeter * (semiPerimeter - a) * (semiPerimeter - b) * (semiPerimeter - c))
			log.Printf("Heronian area is:%f", area)
		} else {
			log.Println("Invalid heronian")
		}
	} else {
		log.Println("Invlid triangel")
	}
}

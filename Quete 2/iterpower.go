package main

import (
	"fmt"
	"math"
)

func IterativePowerl(nb int, power int) {
	if nb > 0 {
		fmt.Println(math.Pow(float64(nb), float64(power)))
	} else if power <= 0 {
		fmt.Println("0")
	}
}

func main() {
	IterativePowerl(7, 3)
}

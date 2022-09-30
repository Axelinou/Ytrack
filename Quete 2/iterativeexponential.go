package main

import (
	"fmt"
	"math"
)

func IterativeFactorial(nb int) int {
	if nb >= 0 {
		g := float64(nb)

		y := math.Exp(g)
		fmt.Print(g, "'s exponential value is ", y)
		return (nb)
	}
	return (nb)
}

func main() {
	IterativeFactorial(0)
}

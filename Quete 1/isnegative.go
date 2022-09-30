package main

import (
	"math"

	"github.com/01-edu/z01"
)

func IsNegative(nb int) {

	res_1 := math.Signbit(6)
	if res_1 == true {
		z01.PrintRune('T')
	} else {
		z01.PrintRune('F')

	}

}
func main() {
	IsNegative(6)
}

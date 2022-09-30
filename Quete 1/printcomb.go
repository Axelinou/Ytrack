package main

import (
	"github.com/01-edu/z01"
)

func Comb() {
	for x := '0'; x != '9'+3; x++ {
		for y := '0'; y != '9'+1; y++ {
			for z := '0'; z != '9'+1; z++ {
				if x < y {
					if y < z {
						z01.PrintRune(x)
						z01.PrintRune(y)
						z01.PrintRune(z)
						z01.PrintRune(',')

					}
				}
			}
		}

	}

}

func main() {
	Comb()
}

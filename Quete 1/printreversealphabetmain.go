package main

import "github.com/01-edu/z01"

func main() {
	for lettre := 'a'; lettre != 'z'+1; lettre++ {
		z01.PrintRune(lettre)
	}
}

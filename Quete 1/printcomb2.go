package main

import "github.com/01-edu/z01"

func PrintComb2() {
	for z := '0'; z <= '9'; z++ {
		for y := '0'; y <= '9'; y++ {
			for i := '0'; i <= '9'; i++ {
				for j := y + 1; j <= '9'; j++ {

					z01.PrintRune(z)
					z01.PrintRune(y)
					z01.PrintRune(' ')
					z01.PrintRune(i)
					z01.PrintRune(j)

					if !(z == '9' && y == '8' && i == '9' && j == '9') {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
}

func main() {
	PrintComb2()
}

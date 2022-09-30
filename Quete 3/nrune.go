package main

import "fmt"

func NRune(texte string, n int) rune {
	r := []rune(texte)
	fmt.Println(r)
	if n-1 < 0 || n > len(r) {
		return 0
	}
	return r[n-1]
}
func main() {
	fmt.Println(NRune("abcd", 3))
}

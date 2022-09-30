package main

import "fmt"

func AlphaCount(chaine string) int {
	var count int = 0
	for i := 0; i < len(chaine); i++ {
		if chaine[i] >= 'a' && chaine[i] <= 'z' || chaine[i] >= 'A' && chaine[i] <= 'Z' {
			count = count + 1
		}
	}
	return count
}

func main() {
	s := "Hello 78 World!    4455 /"
	nb := AlphaCount(s)
	fmt.Println(nb)
}

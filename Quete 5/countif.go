package main

import "fmt"

func main() {
	tab1 := []string{"Hello", "how", "are", "you"}
	tab2 := []string{"This", "1", "is", "4", "you"}
	answer1 := CountIf(IsNumeric, tab1)
	answer2 := CountIf(IsNumeric, tab2)
	fmt.Println(answer1)
	fmt.Println(answer2)
}

func CountIf(f func(string) bool, a []string) int {
	var count int
	for _, str := range a {
		if f(str) {
			count++
		}
	}
	return count
}

func IsNumeric(s string) bool {
	if s == "" {
		return false
	}

	for i := 0; i < len(s)-1; i++ {
		if s[i] < 48 || s[i] > 57 {
			return false
		}
	}
	return true

}

//Countif

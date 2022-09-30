package main

import "fmt"

func IterativeFactorial(nb int) int {
	var result int = 0
	var absoluteNb int = 0
	if nb < -20 || nb > 20 {
		return 0
	} else if nb < 0 {
		absoluteNb = -nb
		result = -1
	} else if nb > 0 {
		absoluteNb = nb
		result = 1
	} else if nb == 0 {
		return 1
	}
	if nb == 0 {
		result = 0
	} else {
		for i := 1; i <= absoluteNb; i++ {
			result *= i
		}
	}
	return result
}
func main() {
	fmt.Println(IterativeFactorial(5))
}

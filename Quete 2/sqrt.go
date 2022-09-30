package main

import "fmt"

func Sqrt(nb int) int {
	var a int = 0
	for i := 1; i <= nb; i++ {
		if i*i == nb {
			a = i
			return a
		}
	}
	return 0
}

func main() {
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(3))
}

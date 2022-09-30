package main

import "fmt"

func MakeRange(min, max int) []int {
	var plain int
	if min < max {
		var tab []int
		for i := min; i < max; i++ {
			plain += tab[]
		}
		return tab

	} else {
		return nil
	}
}

func main() {
	fmt.Println(MakeRange(5, 10))
}

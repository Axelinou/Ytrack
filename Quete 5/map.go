package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	result := Map(IsPrime, a)
	fmt.Println(result)
}

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func Map(f func(int) bool, a []int) []bool {
	var tableauprime []bool

	for _, tgv := range a {
		tableauprime = append(tableauprime, f(tgv))

	}
	return tableauprime
}

//map

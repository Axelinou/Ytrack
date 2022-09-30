package main

import "fmt"

func FindNextPrime(nb int) int {
	for i := nb; true; i++ {
		if IsPrime(i) {
			return i
		}
	}
	return 0
}

func IsPrime(nb int) bool {
	if nb <= 3 {
		return nb > 1
	}
	if nb%2 == 0 || nb%3 == 0 {
		return false
	}
	i := 5
	for i*i <= nb {
		if nb%i == 0 || nb%(i+2) == 0 {
			return false
		}
		i += 6
	}
	return true
}

func main() {
	fmt.Println(FindNextPrime(5))
	fmt.Println(FindNextPrime(4))
}

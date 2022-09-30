package main

import "fmt"

func main() {
	a := 0
	b := &a
	n := &b
	UltimatePointOne(&n)
	fmt.Println(a)
}

func UltimatePointOne(n ***int) {
	***n = 1
}

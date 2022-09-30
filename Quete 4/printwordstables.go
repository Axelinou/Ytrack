package main

import "fmt"

func PrintWordsTables(table []string) {
	for i := 0; i < len(table); i++ {
		fmt.Println(table[i])
	}
}
func main() {
	test := []string{"Hello", "how", "are", "you?"}
	PrintWordsTables(test)
}

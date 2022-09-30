package main

import (
	"fmt"
)

func ConcatParams(args []string) string {
	var str string
	for i := 0; i < len(args); i++ {
		str += args[i] + "\n"
	}
	return str
}
func main() {
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(ConcatParams(test))
}

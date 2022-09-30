package main

import "fmt"

func splitwhitespaces(str string) []string {
	var tab []string
	for _, word := range str {
		if word == "d" {
			tab = append(tab, word)
		}
	}
	return tab

}

func main() {
	fmt.Printf("%s", splitwhitespaces("Hello how are you?"))
}

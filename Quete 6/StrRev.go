package main 

import "fmt"



func main() {
	s := "Hello World!"
	s = StrRev(s)
	fmt.Println(s)
}

func StrRev(s string) string {
	chars := []rune(s)
	for i, s := 0, len(chars)-1; i < s; i, s = i+1, s-1 {
		chars[i], chars[s] = chars[s], chars[i]
	}
	return string(chars)
}

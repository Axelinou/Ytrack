package main 

import "fmt"

func main() {
	a := 0
	b := 1
	Swap(&a, &b)
	fmt.Println(a)
	fmt.Println(b)

}

func Swap(a *int, b *int) {
	fmt.Println(a, b) //0, 1
	a, b = b, a
	fmt.Println(a, b) //1, 0
}

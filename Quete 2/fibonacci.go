package main

func Fibonacci(nb int) int {
	if nb < 0 {
		return -1
	} else if nb <= 1 {
		return nb
	} else {
		return Fibonacci(nb-1) + Fibonacci(nb-2)
	}
}

func main() {
	println(Fibonacci(1))
}

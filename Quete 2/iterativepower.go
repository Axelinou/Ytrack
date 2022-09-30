package main

func IterativePower(nb int, power int) int {
	var result int = nb
	if power == 0 {
		return 1
	} else if power < 0 {
		return 0
	} else {
		for power > 1 {
			power--
			result = nb * result
		}
	}
	return result
}

func main() {
	println(IterativePower(5, 3))
}

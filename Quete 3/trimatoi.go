package main

func TrimAtoi(s string) int {
	index := 0
	sign := 1
	for i := range s {
		if i >= '0' && i <= '9' {
			i = i - '0'
			index = index*10 + int(i)
		} else if index == 0 && i == '-' {
			sign = -1
		}
	}
	return sign * index
}

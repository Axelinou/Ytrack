package main

func foreach(tab []int, f func(int)) {
	for _, v := range tab {
		f(v)
	}
}
func main() {
	tab := []int{1, 2, 3, 4, 5}
	foreach(tab, func(v int) {
		println(v)
	})
}

package main

import "fmt"

func vars() (int, int) {
	return 3, 7
}

func main() {
	a, b := vars()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vars()
	fmt.Println(c)

	d, _ := vars()
	fmt.Println(d)
}

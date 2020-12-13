package main

import "fmt"

func vals() (int, int) {
	retunr 3, 7
}
func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}

package main

import "fmt"

/*
	スライスのゼロ値はnil
	nilスライス:
		長さ = 0
		キャパシティ = 0
*/

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

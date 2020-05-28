package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	// sliceの要素を変更すると、その元となる配列も変更される
	// 同じ元となる配列を共有している他のスライスは、それらの変更が反映される
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

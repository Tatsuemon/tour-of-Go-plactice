package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

// c++と同じように'': char, "": string
func main(){
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
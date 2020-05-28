package main

import "fmt"

// Vertex X:int, Y:int
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}

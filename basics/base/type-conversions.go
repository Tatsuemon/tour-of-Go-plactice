package main

import (
	"fmt"
	"math"
)

func main(){
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	// 型変換
	/*
		i := 42
		f := float64(i)
		u := uint(f)
	*/
	var z unit = uint(f)
	fmt.Println(ln(x, y, z))
}
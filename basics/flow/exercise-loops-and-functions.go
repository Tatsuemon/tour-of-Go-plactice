package main

import (
	"fmt"
	"math"
)

// Sqrt return float64
func Sqrt(x float64) float64 {
	z := float64(5)
	i := 0
	for {
		tmp := float64(z)
		tmp -= (z*z - x) / (2 * z)

		if math.Abs(tmp-z) < 1e-10 {
			break
		}
		z = tmp
		i++
	}
	fmt.Println(i)
	return z
}

func main() {
	myans := Sqrt(2)
	ans := math.Sqrt(2)
	fmt.Println(ans, myans)
	fmt.Println(ans - myans)
}

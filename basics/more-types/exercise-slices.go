package main

import "golang.org/x/tour/pic"

/*
	二次元スライス

	a := make([][]T, i)
	for i := range a {
		a[i] = make([]T, j)
		for j := range a[i] {
			a[i][j] = bb
		}
	}

*/

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)

	for y := range pic {
		pic[y] = make([]uint8, dx)
		for x := range pic[y] {
			pic[y][x] = uint8(x * y)
		}
	}
	return pic

}

func main() {
	pic.Show(Pic)
}

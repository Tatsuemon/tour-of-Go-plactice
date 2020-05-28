package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	/*
		スライスをrangeで繰り返す場合
			i, v := range slice
			iはindex, vは要素のコピー
	*/
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

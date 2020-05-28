package main

import "fmt"

func main() {
	pow := make([]int, 10)
	// indexだけを取得
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	// 要素だけを取得
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

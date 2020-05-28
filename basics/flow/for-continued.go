package main

import "fmt"

func main() {
	sum := 1

	// 初期値とインクリメントは任意　-> 勝手に入る
	// whileはない
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

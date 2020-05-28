package main

import "fmt"

/*
	スライスは長さとキャパシティを持っている
	len(s), cap(s)

	キャパシティは変化しない

	s[low : high]
	スライスのhighのデフォルトはそのスライスの最後の要素まで
	highが参照外になると元の配列からとってくる cap(s)まで

*/

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s) // []

	// Extend its length.
	s = s[:4]
	printSlice(s) // [2, 3, 5, 7]  highの参照がスライスの参照外になっているため、元の配列を参照している

	// Drop its first two values.
	s = s[2:]
	printSlice(s) // [5, 7]  highのデフォルトはそのスライスの最後の要素まで
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

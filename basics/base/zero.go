package main

import "fmt"

// 初期値を与えずに宣言するとゼロ値が与えられる
func main(){
	var i int // 0
	var f float64 // 0
	var b bool // false
	var s string // ""
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
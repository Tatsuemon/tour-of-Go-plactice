package main

import "fmt"

// 関数の二つ以上の引数が同じ型である場合は最後の方を残して省略して記述できる
func add(x, y int) int {
	return x + y
}

func main(){
	fmt.Println(add(42, 13))
}
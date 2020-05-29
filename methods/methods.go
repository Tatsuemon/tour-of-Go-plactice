package main

import (
	"fmt"
	"math"
)

/*
	Goにはクラスがない
	型にmethodを定義できる

	レシーバ: T

	hogeメソッドはvという名前のT型のレシーバをもつ

	func (v T) hoge() T' {

	}
*/

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

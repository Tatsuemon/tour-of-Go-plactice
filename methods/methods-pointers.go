package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*
	レシーバ自体(ここでいうVertex)を変更する時に、ポインタレシーバを使用する

	変数レシーバ（今までのメソッドでの書き方）: メソッド内では変数のコピーが扱われる
	ポインタレシーバ（レシーバにポインタを設定する書き方）: オリジナルを参照している
*/

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}

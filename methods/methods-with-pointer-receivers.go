package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*
	ポインタレシーバを使用する理由
	1. メソッドがレシーバが指す先の変数を変更するため
	2. メソッドの呼び出し毎に変数のコピーを避けるため

	一般的には、変数レシーバ、または、ポインタレシーバのどちらかですべてのメソッドを与え、混在させるべきではない
*/

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}

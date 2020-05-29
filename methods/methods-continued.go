package main

import (
	"fmt"
	"math"
)

type MyFloat float64

/*
	structだけでなく, typeにもメソッドを宣言できる

	レシーバを伴うメソッドの宣言は、レシーバ型が同じパッケージにある必要があります。
	他のパッケージに定義している型に対して、レシーバを伴うメソッドを宣言できません （組み込みの int などの型も同様です）。
*/

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

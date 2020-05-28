package main

import (
	"fmt"
	"math/cmplx"
)
/*
	bool

	string

	int int8 int16 int32 int64
	uint uint8 uint16 uint32 uint64 uintptr

	byte // uint8の別名

	rune // int32の別名, Unicodeのコードポイントを表す

	float32 float64

	complex64 complex128
*/

var (
	ToBe bool = false
	MaxInt uint64 = 1<<64 - 1
	z complex128 = cmplx.Sqrt(-1 + 12i)
)

func main(){
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
package main

import "fmt"

func main(){
	/*
		空のインターフェースは任意の方を保持できる
	*/
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
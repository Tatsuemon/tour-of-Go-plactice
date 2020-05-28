package main

import "fmt"

/*
	名前をつけた戻り値の変数を使うと、return ステートメントに何も書かずに戻すことができる
	短い関数で使うべき
*/
func split(sum int) (x, y int){
	x = sum*4/9
	y = sum-x
	return 
}

func main(){
	fmt.Println(split(17))
}
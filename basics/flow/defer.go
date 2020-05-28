package main

import "fmt"

func main() {
	/*
		defer で渡した関数の実行を呼び出し元の関数の終わりまで遅延させる

		引数はすぐに評価されるが、関数自体は呼び出し元（ここではmain）がreturnするまで実行されない
	*/
	defer fmt.Println("world")

	fmt.Println("hello")
}

package main

import "fmt"

func main() {
	// バッファ, 第二引数はバッファの長さ
	/*
		バッファが詰まったときはチャネルへの送信をブロックする
		バッファがからのときはチャネルの受信をブロックする
	*/
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// チャネルのclose
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)

	/*
		v, ok <- ch

		okにチャネルが開いているかが出る
	*/

	// チャネルが閉じるまでチャネルから値を繰り返し受信し続ける
	for i := range c {
		fmt.Println(i)
	}
}

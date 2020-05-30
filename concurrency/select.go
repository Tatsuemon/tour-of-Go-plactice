package main

import "fmt"

func fibonacci(c, quit chan int){
	x, y := 0, 1

	for {
		/*
			準備ができたcaseをランダムで実行する
		*/
		select {
		case c<-x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main(){
	c := make(chan int)
	quit := make(chan int)

	// 即時関数
	go func(){
		for i:=0; i<10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacci(c, quit)

}
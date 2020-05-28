package main 

import "fmt"

// 定数
/*
	char, string bool, numericのみで使える
	定数は := で宣言できない
*/
const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
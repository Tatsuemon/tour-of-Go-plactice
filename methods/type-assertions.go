package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	// s, ok := i.(T) でインターフェースが型Tを保持しているかがわかる
	s, ok := i.(string)
	fmt.Println(s, ok)

	// これでnil判定的なことができる
	if ok {
		fmt.Println(s)
	}

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}

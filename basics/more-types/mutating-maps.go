package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// i, j := m['key'] -> i: value, j: 存在するかどうか
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	// c++と同じ挙動
	m["Answer"]++
	fmt.Println("The value:", m["Answer"])
}

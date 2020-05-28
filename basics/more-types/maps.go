package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// map[T]T' ->  key: T, value: T'
// mapのゼロ値はnil, nilマップの場合はキーを持たず、キーの追加もできない
// makeで使用できるようにしている
var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)

	fmt.Println(m)

	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
	fmt.Println(m)
}

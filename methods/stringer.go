package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}

	// Stringが勝手に呼ばれる これがStringer??
	/*
		type Stringer interface {
			String() string
		}
	*/
	fmt.Println(a, z)
}

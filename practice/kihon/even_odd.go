package main

import "fmt"

func useif(){
	for i:=1; i<=100; i++ {
		if i%2==1 {
			fmt.Printf("%d-奇数\n", i)
		}else{
			fmt.Printf("%d-偶数\n", i)
		}
	}
}

func useswitch(){
	for i:=1; i<=100; i++ {
		switch {
		case i%2==1:
			fmt.Printf("%d-奇数\n", i)
		default:
			fmt.Printf("%d-偶数\n", i)
		}
	}
	return
}

func main(){
	// useif()
	useswitch()
}
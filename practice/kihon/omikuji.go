package main

import (
	"fmt"
	"math/rand"
	"time"
)

func omikuji(n int) string {
	switch n {
	case 1:
		return fmt.Sprint("凶")
	case 2, 3:
		return fmt.Sprint("吉")
	case 4, 5:
		return fmt.Sprint("中吉")
	case 6:
		return fmt.Sprint("大吉")
	default:
		return fmt.Sprint("何もないで")
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(6)
	fmt.Println(i)
	fmt.Println(omikuji(i + 1))
}

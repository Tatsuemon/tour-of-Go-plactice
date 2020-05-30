package main

func main() {
	ns := [4]int{19, 86, 1, 12}
	var sum int
	for _, v := range ns {
		sum += v
	}
	println(sum)
}

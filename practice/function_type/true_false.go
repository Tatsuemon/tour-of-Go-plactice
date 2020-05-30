package main

func main() {
	m := [5]int{1, 2, 3, 4, 5}
	// println(len(m))

	s := m[:5]
	// println(len(s))
	// pp(s)

	s = s[:4]
	println(len(s))
	pp(s)

	s[3] = 10

	s = s[:5]
	println(len(s))
	pp(s)

	print(m[3])

}

func pp(a []int) {
	for _, v := range a {
		println(v)
	}
}

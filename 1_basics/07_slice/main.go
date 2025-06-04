package main

import ("fmt")

func main() {
	var s[] string // nil, lengh = 0
	s = make([]string, 3) // length = 3
	s[0], s[1], s[2] = "a", "b", "c"
	fmt.Println(s)

	s = append(s, "d", "e")
	fmt.Println(s)
	c := make([]string, len(s))
	copy(c, s)
	
	l := s[2:5] // s[2], s[3], s[4]
	fmt.Println(l)

	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}

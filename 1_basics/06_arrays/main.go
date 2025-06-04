package main

import ("fmt")

func main() {
	// var arr = [len]datatype{values}
	// or, var arr = [...]datatype{values}
	// arr := [len]dt{val}
	// or, arr := [...]dt{val}
	
	var arr = [3]int{1, 2, 3}
	fmt.Println(arr)
	// arr[0] = 1

	// arr := [5]int{} // [0 0 0 0 0]
	// arr := [5]int{1, 2} // [1 2 0 0 0]

	// initialize only specific elements
	// arr := [5]int{1:10, 2:40}

	// len(arr) = size of array
	
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	b = [...]int{5, 4, 3, 2, 1}
	fmt.Println(b)
	
	b = [...]int{100, 3:400, 500}
	fmt.Println(b)

	var twoD [2][3] int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i+j
		}
	}
	fmt.Println(twoD)
  twoD = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(twoD)
}

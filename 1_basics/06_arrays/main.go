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
}

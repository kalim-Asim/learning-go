package main

import "fmt" 

func printSlice(items []int) {
	for _, item := range items {
		fmt.Print(item)
		fmt.Print(" ")
	}
	fmt.Println()
}

// generic
func print[T int | string](items []T) {
	for _, item := range items {
		fmt.Print(item)
		fmt.Print(" ")
	}
	fmt.Println()
}

func main() {
	printSlice([]int{1, 2, 3})
	
	print([]int{2,4,6})
	print([]string{"abc", "cde"})
}

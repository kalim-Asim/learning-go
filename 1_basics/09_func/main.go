package main

import ("fmt")

//simple function
func plus(a int, b int) int {
	return a + b
}

// multiple return values
func fun()(int, int) {
	return 1, 2
}

// variadic function
// can be called with any no. of trailing arg
func sum(nums ...int) {
	tot := 0
	for _, num := range nums {
		tot += num
	}
	fmt.Println(tot)
}

func f() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func main() {
	a, b := fun()
	fmt.Println(a, b)

	sum(1,2)
	sum(1,2,3)

	// recursive function
	var fib func(n int) int
	fib = func(n int) int{
		if n < 2 {return n}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(7))

	/* CLOSURES */
	f2 := f()
	fmt.Println(f2()) // 1
	fmt.Println(f2()) // 2

	f3 := f()
	fmt.Println(f3()) // 1
}

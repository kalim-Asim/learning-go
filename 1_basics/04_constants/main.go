package main

import "fmt"

const age int = 50
//var name = "golang" (works)
//name := "golang" (not works)

func main() {
	//fmt.Println(age)
	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port, host)
}


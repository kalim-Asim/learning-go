package main

import "fmt"

// interface
type geometry interface {
	area() int
	perimeter() int
}

type rect struct {
	width, height int
}

func (r rect) perimeter() int {
	return 2*r.width + 2*r.height
}

func (r *rect) area() int {
	return r.width * r.height
}

func main() {
	 r := rect{width: 5, height: 10}

	 fmt.Println(r.area())
	 fmt.Println(r.perimeter())
}

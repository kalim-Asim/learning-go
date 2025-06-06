package main

import "fmt"

type OrderStatus int

const (
	Received OrderStatus = iota // 0
	Confirmed // 1
	Prepared  // 2 
	Delivered // 3
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Changing order status to", status)
}

func main() {
	changeOrderStatus(Delivered) // 3	
}

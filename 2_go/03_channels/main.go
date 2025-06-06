package main

import (
	"fmt"
	"time"
)

// func processNum(numChan chan int) {
 	// fmt.Println("Processing Number", <-numChan)

// 	for num := range numChan {
// 		fmt.Println("processing no.", num)
// 		time.Sleep(time.Second)
// 	}
// }


func sum(result chan int, a int, b int) {
	time.Sleep(time.Second * 2)
	result <- a + b
	fmt.Println(result)
}

func main() {
	/*
		    messageChan := make(chan string)

				//send
				messageChan <- "ping"//blocking

				//receive
		  	msg := <- messageChan
				fmt.Println(msg)

	*/

	/*
		numChan := make(chan int)

		go processNum(numChan)

		// numChan <- 5
		for {
			numChan <- rand.Intn(100)
		}
	*/

	result := make(chan int)
	go sum(result, 2, 3)
	res := <-result
	fmt.Println(res)
}
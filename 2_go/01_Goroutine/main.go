package main

import "fmt"
import "time"

func task(i int) {
	fmt.Println(i)
}

func main() {
	for i := 0; i <= 10; i++ {
		// go task(i)

		go func(i int) {
			fmt.Println(i)
		}(i)

	}
	// if not written goroutine gives no o/p
	time.Sleep(time.Second * 1)
}

package main

import (
	"fmt"
	"sync"
)

func task(id int, w* sync.WaitGroup) {
	defer w.Done() // after function is done, defer runs
	fmt.Println("doing task : ", id)
}

func main() {
	var wg sync.WaitGroup
	// Add -> done -> wait

	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go task(i, &wg)
	}
	
	wg.Wait()
}


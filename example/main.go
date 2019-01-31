package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i <= 9; i++ {

		go func() {
			fmt.Println("", i)
			wg.Done()
		}()
	}
	wg.Wait()
}

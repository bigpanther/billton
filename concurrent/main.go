package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		go func(i int) {
			fmt.Println(i)
			fmt.Println("some text")
			wg.Done()
		}(i)
		wg.Add(1)
	}
	wg.Wait()
}

package main

import (
	"fmt"
	"sync"
)

func lock() {
	var mu sync.Mutex
	wg := sync.WaitGroup{}
	var a = 0
	for i := 0; i < 2000; i++ {
		go func(i int) {
			mu.Lock()
			var y = a
			fmt.Println("hah")
			a = y + 1
			mu.Unlock()
			// fmt.Println(i)
			// fmt.Println("some text")
			wg.Done()
		}(i)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(a)
}

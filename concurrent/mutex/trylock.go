package main

import (
	"fmt"
	"sync"
)

func tryLock() {
	var mu sync.Mutex
	wg := sync.WaitGroup{}
	var a = 0
	for i := 0; i < 2000; i++ {
		go func(i int) {
			for {
				if mu.TryLock() {
					var y = a
					fmt.Println("hah")
					a = y + 1
					mu.Unlock()
					break
				} else {
					fmt.Println("do grocery")
				}
			}
			wg.Done()
		}(i)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(a)
}

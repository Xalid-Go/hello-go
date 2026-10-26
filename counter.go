package main

import "sync"

// Count запускает n горутин, каждая увеличивает общий счётчик.
func Count(n int) int {
	var wg sync.WaitGroup
	total := 0
	for range n {
		wg.Add(1)
		go func() {
			defer wg.Done()
			total++
		}()
	}
	wg.Wait()
	return total
}

package main

import (
	"sync"
	"sync/atomic"
)

func Count(n int) int {
	var wg sync.WaitGroup
	var total atomic.Int64

	for range n {
		wg.Add(1)
		go func() {
			defer wg.Done()
			total.Add(1)
		}()
	}
	wg.Wait()
	return int(total.Load())
}

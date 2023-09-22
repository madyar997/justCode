package main

import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[string]int)

	wg := sync.WaitGroup{}

	mu := sync.Mutex{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			m["some_key"]++
			mu.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println(m["some_key"])
}

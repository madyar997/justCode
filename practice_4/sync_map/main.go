package main

import (
	"fmt"
	"sync"
)

func main() {
	m := sync.Map{}

	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {

		wg.Add(1)

		go func() {
			defer wg.Done()
			value, ok := m.Load("some_key")
			if ok {
				m.Store("some_key", value.(int)+1)
			} else {
				m.Store("some_key", 1)
			}
		}()
	}

	wg.Wait()

	fmt.Println(m.Load("some_key"))
}

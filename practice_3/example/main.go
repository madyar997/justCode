package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	wg := &sync.WaitGroup{}

	wg.Add(3)

	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second) //имитация бурной деятельности
		fmt.Println("first goroutine finished")
	}()

	go func() {
		defer wg.Done()
		//time.Sleep(1 * time.Second)
		fmt.Println("second goroutine finished")
	}()

	go func() {
		defer wg.Done()
		//time.Sleep(1 * time.Second)
		fmt.Println("third goroutine finished")
	}()

	wg.Wait()

	fmt.Println("All goroutines are finsihed their work. ")
}

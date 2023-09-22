package main

import (
	"fmt"
	"time"
)

func main1() {
	ticker := time.NewTicker(5 * time.Second)

	go func() {
		for t := range ticker.C {
			fmt.Println("tick at ", t)
		}
	}()

	time.Sleep(15 * time.Second)
	ticker.Stop()
}

func main2() {
	ticker := time.NewTicker(5 * time.Second)

	go func() {
		for t := time.Now(); ; t = <-ticker.C {
			fmt.Println("tick at ", t)
		}
	}()

	time.Sleep(15 * time.Second)
	ticker.Stop()
}

func main3() {
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan struct{})

	go func() {
		for {
			select {
			case t := <-ticker.C:
				fmt.Println("tick at ", t)
			case <-done:
				return
			}
		}

	}()

	time.Sleep(5 * time.Second)
	ticker.Stop()

	done <- struct{}{}
	fmt.Println("done")
}

package main

import "fmt"

func main() {
	numJobs := 100
	numWorker := 5

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for i := 1; i <= numWorker; i++ {
		go worker(square, jobs, results, i)
	}

	for i := 1; i <= numJobs; i++ {
		jobs <- 5
	}

	close(jobs)

	for i := 0; i < numJobs; i++ {
		fmt.Println(<-results)
	}
}

func worker(f func(int) int, jobs <-chan int, results chan<- int, id int) {
	for num := range jobs {
		fmt.Printf("worker %d is in progress \n", id)
		results <- f(num)
	}
}

func square(num int) int {
	return num * num
}

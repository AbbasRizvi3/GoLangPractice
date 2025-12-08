package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func makeJobs(x int, job chan int) {
	wg.Add(1)
	job <- x

}

func doJob(job <-chan int, result chan<- int) {
	for j := range job {
		ans := j * 2
		result <- ans
		wg.Done()
	}
}

func main() {
	numJobs := 5
	numWorkers := 3

	jobs := make(chan int)
	results := make(chan int)

	go func() {
		for val := range results {
			fmt.Println("Result received:", val)
		}
	}()

	for i := 0; i < numWorkers; i++ {
		go doJob(jobs, results)
	}

	for i := 0; i < numJobs; i++ {
		makeJobs(i, jobs)
	}

	wg.Wait()
	close(jobs)
	close(results)
}

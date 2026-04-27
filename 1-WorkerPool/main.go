package main

import (
	"fmt"
	"sync"
	"time"
)

func RunWorkers(numWorkers int, numJobs int) {

	jobs := make(chan int, numJobs)
	var wg sync.WaitGroup

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			for job := range jobs {
				fmt.Printf("Worker %d processing job %d\n", id, job)
				time.Sleep(1 * time.Second)
			}
		}(w)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	close(jobs)
	wg.Wait()
}

func main() {
	RunWorkers(2, 10)
}

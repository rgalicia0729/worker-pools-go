package main

import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func workers(jobs <-chan int, fib chan<- int) {
	for job := range jobs {
		resultFib := fibonacci(job)
		fib <- resultFib
	}
}

func main() {
	tasks := []int{3, 4, 5, 8, 13, 40, 20, 25, 10}
	nWorkers := 3

	jobs := make(chan int, len(tasks))
	fib := make(chan int, len(tasks))

	for i := 0; i < nWorkers; i++ {
		go workers(jobs, fib)
	}

	for _, value := range tasks {
		jobs <- value
	}
	close(jobs)

	for i := 0; i < len(tasks); i++ {
		fmt.Println(<-fib)
	}
}

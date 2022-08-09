package main

import (
	"fmt"
	"time"
)

func main() {
	tasks := gen(1, 2, 3, 4, 5)
	limitter := time.Tick(time.Second * 2)

	results := processorWithLimit(tasks, limitter)

	for result := range results {
		fmt.Println("Result: ", result)
	}
}

func gen(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}

func processorWithLimit(in <-chan int, limitter <-chan time.Time) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for val := range in {
			select {
			case out <- val * 2:
				time.Sleep(time.Second)
			case <-limitter:
				fmt.Println("wait...")
				time.Sleep(time.Second * 2)
				out <- val * 2
			}
		}
	}()

	return out
}

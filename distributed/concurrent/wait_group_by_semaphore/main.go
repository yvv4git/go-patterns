package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	n := len(numbers)

	sem := New(n)

	for _, num := range numbers {
		go func(n int) {
			fmt.Println(n)
			sem.Inc(1)
		}(num)
	}

	sem.Dec(n)

}

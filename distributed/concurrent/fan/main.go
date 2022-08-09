package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen(2, 3)

	// Distribute the sq work across two goroutines that both read from in.
	// FAN-OUT
	c1 := sq(in)
	c2 := sq(in)

	// Consume the merged output from c1 and c2.
	// FAN-IN
	for n := range merge(c1, c2) {
		fmt.Println(n) // 4 then 9, or 9 then 4
	}
}

// Step-3: объединяем результаты
func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// Step-2: Читаем из канала задач и выполняем полезную работы, результат возвращаем в выходной канал
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// Step-1: генератор конвертирует входные данные в канал, генерит значения
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

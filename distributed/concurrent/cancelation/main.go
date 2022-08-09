package main

import (
	"fmt"
	"time"
)

func main() {
	chDone := make(chan struct{})

	go func() {
		time.Sleep(time.Second * 2)
		chDone <- struct{}{} // прекращаем выполнение задач
	}()

	in := gen(1, 2, 3, 4, 5)
	out := sq(chDone, in)

	for result := range out {
		fmt.Println("RESULT: ", result)
	}
}

// Step-2: обработчик задач с возможностью явно прекратить выполнение
func sq(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for n := range in {
			time.Sleep(time.Second) // Симулируем время на работу

			select {
			case out <- n * n:
			case <-done:
				return
			}
		}
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

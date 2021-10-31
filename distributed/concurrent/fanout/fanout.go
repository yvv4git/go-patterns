package fanout

// FanOut - used for de-multiplexing goroutines
func FanOut(source <-chan int, n int) []<-chan int {
	dest := make([]<-chan int, 0)

	for i := 0; i < n; i++ {
		ch := make(chan int)
		dest = append(dest, ch)

		go func() {
			defer close(ch)
			for val := range source {
				ch <- val
			}
		}()
	}

	return dest
}

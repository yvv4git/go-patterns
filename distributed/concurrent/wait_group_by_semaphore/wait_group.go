package main

// Semaphore - семафор - канал из пустых структур
type Semaphore chan struct{}

func New(n int) Semaphore {
	return make(Semaphore, n)
}

// Inc - у wait-group есть метод add, который увеличивает значение.
func (s Semaphore) Inc(k int) {
	for i := 0; i < k; i++ {
		s <- struct{}{}
	}
}

func (s Semaphore) Dec(k int) {
	for i := 0; i < k; i++ {
		<-s
	}
}

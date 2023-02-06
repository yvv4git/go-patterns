package worker_pool

import (
	"fmt"
)

// Worker контролирует всю работу
type Worker struct {
	ID       int
	taskChan chan *Task
	quit     chan bool
}

// NewWorker возвращает новый экземпляр worker-а
func NewWorker(channel chan *Task, ID int) *Worker {
	return &Worker{
		ID:       ID,
		taskChan: channel,
		quit:     make(chan bool),
	}
}

// StartBackground запускает worker-а в фоне
func (wr *Worker) StartBackground() {
	fmt.Printf("Starting worker %d\n", wr.ID)

	for {
		select {
		case task := <-wr.taskChan:
			process(wr.ID, task)
		case <-wr.quit:
			return
		}
	}
}

// Stop Остановка quits для воркера
func (wr *Worker) Stop() {
	fmt.Printf("Closing worker %d\n", wr.ID)
	go func() {
		wr.quit <- true
	}()
}

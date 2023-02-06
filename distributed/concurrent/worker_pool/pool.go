package worker_pool

import (
	"fmt"
	"sync"
	"time"
)

type Pool struct {
	Tasks   []*Task
	Workers []*Worker

	concurrency   int
	collector     chan *Task
	runBackground chan bool
	wg            sync.WaitGroup
}

// AddTask добавляет таски в pool
func (p *Pool) AddTask(task *Task) {
	p.collector <- task
}

// RunBackground запускает pool в фоне
func (p *Pool) RunBackground() {
	go func() {
		for {
			fmt.Print("Waiting for tasks to come in ...\n")
			time.Sleep(10 * time.Second)
		}
	}()

	for i := 1; i <= p.concurrency; i++ {
		worker := NewWorker(p.collector, i)
		p.Workers = append(p.Workers, worker)
		go worker.StartBackground()
	}

	for i := range p.Tasks {
		p.collector <- p.Tasks[i]
	}

	p.runBackground = make(chan bool)
	<-p.runBackground
}

// Stop останавливает запущенных в фоне worker-ов
func (p *Pool) Stop() {
	for i := range p.Workers {
		p.Workers[i].Stop()
	}
	p.runBackground <- true
}

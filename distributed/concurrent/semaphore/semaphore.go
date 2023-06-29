package semaphore

import (
	"errors"
	"time"
)

var (
	ErrNoTasks        = errors.New("can't capture the semaphore")
	ErrIllegalRelease = errors.New("can't release the semaphore without capturing it first")
)

type Interface interface {
	Acquire() error
	Release() error
}

type implementation struct {
	sem     chan struct{}
	timeout time.Duration
}

// Acquire - захватить семафор, берем 1 тикет.
func (s *implementation) Acquire() error {
	select {
	case s.sem <- struct{}{}:
		return nil
	case <-time.After(s.timeout):
		return ErrNoTasks
	}
}

// Release - освободить семафор. Минус 1 ticket.
// Соответственно, если тикеты кончились, то захватить семафор не удастся.
func (s *implementation) Release() error {
	select {
	case _ = <-s.sem:
		return nil
	case <-time.After(s.timeout):
		return ErrIllegalRelease
	}
}

func New(tickets int, timeout time.Duration) Interface {
	return &implementation{
		sem:     make(chan struct{}, tickets),
		timeout: timeout,
	}
}

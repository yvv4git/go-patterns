package breaker

import (
	"golang.org/x/net/context"
	"sync"
	"time"
)

// Result - used as result of action
type Result struct {
	X int
}

// Action - some function to be performed
type Action func(ctx context.Context) (Result, error)

// Breaker - some realisation of Breaker pattern
func Breaker(act Action, failureThreshold uint, delay time.Duration) Action {
	var consecutiveFailure int = 0
	var lastAttempt = time.Now()
	var m sync.RWMutex

	// Замыкание подсчитывает количество ошибок и когда число достигает порогового значения, то возвращаем err=nil.
	// Если получим положительный ответ, то счетчик сбрасывается.
	return func(ctx context.Context) (Result, error) {
		m.RLock() // Брем блокировку чтения

		d := consecutiveFailure - int(failureThreshold)
		if d >= 0 {
			shouldRetryAt := lastAttempt.Add(delay << d)
			if !time.Now().After(shouldRetryAt) {
				m.RUnlock()
				return Result{}, nil
			}
		}

		m.RUnlock() // Освобождаем блокировку чтения
		response, err := act(ctx)

		m.Lock() // Блокируемся
		defer m.Unlock()

		lastAttempt = time.Now()
		if err != nil {
			consecutiveFailure++
			return response, err
		}

		consecutiveFailure = 0
		return response, nil
	}
}

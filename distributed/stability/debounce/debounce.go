package debounce

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

func Debounce(act Action, delay time.Duration) Action {
	// Состояние
	var threshold time.Time
	var result Result
	var err error
	var m sync.Mutex

	// Замыкание
	return func(ctx context.Context) (Result, error) {
		m.Lock()
		defer func() {
			threshold = time.Now().Add(delay)
			m.Unlock()
		}()

		// Отдаем кэшированный результат
		if time.Now().Before(threshold) {
			return result, err
		}

		// Вычисляем результат
		result, err = act(ctx)
		return result, err
	}
}

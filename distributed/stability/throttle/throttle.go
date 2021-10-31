package throttle

import (
	"errors"
	"golang.org/x/net/context"
	"sync"
	"time"
)

var (
	ErrToManyCalls   = errors.New("to many calls")
	ErrTimeoutAction = errors.New("timeout error")
)

// Result - used as result of action
type Result struct {
	X int
}

// Action - some function to be performed
type Action func(ctx context.Context) (Result, error)

// Throttle - some implementation of the throttle
func Throttle(act Action, maxAttempts uint, supply uint, delay time.Duration) Action {
	var attempts = maxAttempts
	var once sync.Once

	return func(ctx context.Context) (Result, error) {
		if ctx.Err() != nil {
			return Result{}, ctx.Err()
		}

		once.Do(func() {
			ticker := time.NewTicker(delay)
			go func() {
				defer ticker.Stop()

				for {
					select {
					case <-ctx.Done():
						return
					case <-ticker.C:
						a := attempts + supply
						if a > maxAttempts {
							a = maxAttempts
						}
						attempts = a
					}
				}
			}()
		})

		if attempts <= 0 {
			return Result{}, ErrToManyCalls
		}

		attempts--
		return act(ctx)
	}
}

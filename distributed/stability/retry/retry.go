package retry

import (
	"context"
	"time"
)

// Result - used as result of action
type Result struct {
	X int
}

// Action - some function to be performed
type Action func(ctx context.Context) (Result, error)

func Retry(act Action, retries int, delay time.Duration) Action {
	return func(ctx context.Context) (Result, error) {
		for r := 0; ; r++ {
			response, err := act(ctx)
			if err == nil || r >= retries {
				return response, err
			}

			select {
			case <-time.After(delay): // Определяет задержку между попытками
			case <-ctx.Done():
				return Result{}, ctx.Err()
			}
		}
	}
}

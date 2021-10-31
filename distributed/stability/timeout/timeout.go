package timeout

import "golang.org/x/net/context"

type (
	Action      func(int) (int, error)
	WithContext func(context.Context, int) (int, error)
)

func Timeout(act Action) WithContext {
	return func(ctx context.Context, arg int) (int, error) {
		chResult := make(chan int)
		chError := make(chan error)

		go func() {
			res, err := act(arg)
			chResult <- res
			chError <- err
		}()

		select {
		case res := <-chResult:
			return res, <-chError
		case <-ctx.Done():
			return 0, ctx.Err()
		}
	}
}

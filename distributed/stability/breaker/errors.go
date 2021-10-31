package breaker

import "errors"

var (
	ErrTimeoutAction = errors.New("timeout error")
)

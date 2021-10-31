package retry

import "golang.org/x/net/context"

// Activator - contract for activator
//go:generate mockery --name=Activator --output=. --filename=./mock.go --inpackage
type Activator interface {
	Activate(ctx context.Context) (Result, error)
}

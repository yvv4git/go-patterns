package timeout

// Activator - contract for activator
//go:generate mockery --name=Activator --output=. --filename=./mock.go --inpackage
type Activator interface {
	Activate(int) (int, error)
}

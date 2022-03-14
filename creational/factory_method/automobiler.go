package factory_method

// Automobiler - used as contract for all automobiles
type Automobiler interface {
	Name() string
	Power() float32
	Speed() float32
}

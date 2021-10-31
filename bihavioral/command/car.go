package command

// Car - contract for all cars
type Car interface {
	Start() Result
	Stop() Result
}

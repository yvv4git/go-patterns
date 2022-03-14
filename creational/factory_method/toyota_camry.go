package factory_method

// ToyotaCamry - specific implementation of the interface Automobiler
type ToyotaCamry struct {
	name  string
	power float32
	speed float32
}

// NewToyotaCamry - used for create instance of ToyotaCamry
func NewToyotaCamry() ToyotaCamry {
	return ToyotaCamry{
		name:  "Toyota Camry",
		power: 230,
		speed: 280,
	}
}

// Name - used for get property of entity
func (a ToyotaCamry) Name() string {
	return a.name
}

// Power - used for get property of entity
func (a ToyotaCamry) Power() float32 {
	return a.power
}

// Power - used for get property of entity
func (a ToyotaCamry) Speed() float32 {
	return a.speed
}

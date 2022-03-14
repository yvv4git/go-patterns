package factory_method

// RenoLogan - specific implementation of the interface Automobiler
type RenoLogan struct {
	name  string
	power float32
	speed float32
}

// NewRenoLogan - used for create instance of RenoLogan
func NewRenoLogan() RenoLogan {
	return RenoLogan{
		name:  "Renault Logan",
		power: 115,
		speed: 220,
	}
}

// Name - used for get property of entity
func (a RenoLogan) Name() string {
	return a.name
}

// Power - used for get property of entity
func (a RenoLogan) Power() float32 {
	return a.power
}

// Power - used for get property of entity
func (a RenoLogan) Speed() float32 {
	return a.speed
}

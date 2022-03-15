package prototype

// Sheep - specific implementation of mutable interface
type Sheep struct {
	name string
}

// NewSheep - simple factory for create instance of Sheep
func NewSheep(name string) Sheep {
	return Sheep{
		name: name,
	}
}

// Who - get name of sheeps
func (s Sheep) Who() string {
	return s.name
}

// Clone - used for clone sheep
func (s Sheep) Clone() Mutable {
	return &Sheep{
		name: s.name + "_Dolly",
	}
}

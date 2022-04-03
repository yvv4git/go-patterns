package options

const (
	defaultFloors         = 2
	defaultHasFireplace   = true
	defaultNoFireplace    = false
	defaultMaterialWood   = "wood"
	defaultMaterialKerpic = "kerpic"
)

// House - used as house implementation
type House struct {
	Material     string
	HasFireplace bool
	Floors       int
}

// NewHouse  used as simple factory for create House instance
func NewHouse(opts ...HouseOption) *House {
	h := &House{
		Material:     defaultMaterialWood,
		HasFireplace: defaultHasFireplace,
		Floors:       defaultFloors,
	}

	for _, opt := range opts {
		opt(h)
	}

	return h
}

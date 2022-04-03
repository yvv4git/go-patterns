package options

const (
	
)

// HouseOptions - used common type for options
type HouseOption func(*House)

// WithMaterialKerpic - ...
func WithMaterialKerpic() HouseOption {
	return func(h *House) {
		h.Material = defaultMaterialKerpic
	}
}

// WithoutFireplace - ...
func WithoutFireplace() HouseOption {
	return func(h *House) {
		h.HasFireplace = false
	}
}

// WithFloors - ...
func WithFloors(floors int) HouseOption {
	return func(h *House) {
		h.Floors = floors
	}
}

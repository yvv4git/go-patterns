package factory_method

// AutomobileType - used for automobiles
type AutomobileType int

const (
	// TypeRenaultLogan - when we want create instance of RenoLogan
	TypeRenaultLogan AutomobileType = iota + 1

	// TypeToyotaCamry - whe we want create instance of ToyotaCamry
	TypeToyotaCamry
)

// AutomobilerFactory - factory method
func AutomobilerFactory(avtoType AutomobileType) Automobiler {
	switch avtoType {
	case TypeRenaultLogan:
		return NewRenoLogan()
	case TypeToyotaCamry:
		return NewToyotaCamry()
	default:
		return NewRenoLogan()
	}
}

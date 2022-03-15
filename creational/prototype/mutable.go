package prototype

// Mutable - a common contract for all cloned types
type Mutable interface {
	Who() string
	Clone() Mutable
}
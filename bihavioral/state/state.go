package state

// State - used as contract for all state realisations
type State interface {
	Write(words string) string
}

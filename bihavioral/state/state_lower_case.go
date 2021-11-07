package state

import "strings"

// LowerCase - concrete state for keyboard
type LowerCase struct{}

// Write - change the text to lowercase
func (l *LowerCase) Write(words string) string {
	return strings.ToLower(words)
}

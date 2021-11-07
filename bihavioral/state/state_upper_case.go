package state

import "strings"

type UpperCase struct{}

// Write - change the text to uppercase
func (l *UpperCase) Write(words string) string {
	return strings.ToUpper(words)
}

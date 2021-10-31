package command

// Command - common interface for commands
type Command interface {
	Execute() Result
}

package command

// Button - used to execute internal logic
type Button struct {
	command Command
}

// Press - used for activate command
func (b *Button) Press() Result {
	return b.command.Execute()
}

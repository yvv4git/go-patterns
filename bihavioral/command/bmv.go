package command

// BMV - used as concrete engine
type BMV struct {
}

// Start - start bmv engine
func (b BMV) Start() Result {
	return "start"
}

// Stop - stop bmv engine
func (b BMV) Stop() Result {
	return "stop"
}

package command

// EngineStopCommand - used as command for stop engine
type EngineStopCommand struct {
	car Car
}

// Execute - used to execute engine stop logic
func (c *EngineStopCommand) Execute() Result {
	return c.car.Stop()
}

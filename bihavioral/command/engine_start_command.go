package command

// EngineStartCommand - used as command for start engine
type EngineStartCommand struct {
	car Car
}

// Execute - used to execute engine start logic
func (c *EngineStartCommand) Execute() Result {
	return c.car.Start()
}

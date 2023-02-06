package decorator

// ChocolateBall - used as decorator.
type ChocolateBall struct {
	iceCream IceCream
}

// NewChocolateBall - used for create instance of ChocolateBall.
func NewChocolateBall(iceCream IceCream) *ChocolateBall {
	return &ChocolateBall{
		iceCream: iceCream,
	}
}

// GetPrice - used for get price.
func (c *ChocolateBall) GetPrice() float32 {
	return c.iceCream.GetPrice() + 50.0
}

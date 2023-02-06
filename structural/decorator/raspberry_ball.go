package decorator

// RaspberryBall - used as decorator.
type RaspberryBall struct {
	iceCream IceCream
}

// NewRaspberryBall - used for create instance of RaspberryBall.
func NewRaspberryBall(iceCream IceCream) *RaspberryBall {
	return &RaspberryBall{
		iceCream: iceCream,
	}
}

// GetPrice - used for get price of ice cream.
func (r *RaspberryBall) GetPrice() float32 {
	return r.iceCream.GetPrice() + 60
}

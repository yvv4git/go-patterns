package mediator

// Client - used as concrete client for chat
type Client struct {
	messageManager Messenger
}

// NewClient - simple factory for create instance of Client
func NewClient(messageManager Messenger) *Client {
	return &Client{
		messageManager: messageManager,
	}
}

func (c *Client) Send(message string) {
	c.messageManager.Receive(message)
}

package builder

// Builder - common contract for all specific implementations
type Builder interface {
	GetMessage() Message
}

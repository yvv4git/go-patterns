package builder

import (
	"errors"
	"strings"

	"github.com/rs/zerolog"
)

var (
	ErrNoValidEmail = errors.New("invalid email")
)

// EmailBuilder - specific builder for emails
type EmailBuilder struct {
	logger zerolog.Logger
	email  Message
}

// NewEmailBuilder - used for create instance of EmailBuilder
func NewEmailBuilder(logger zerolog.Logger) EmailBuilder {
	return EmailBuilder{
		logger: logger,
	}
}

// To - used for set "to" address
func (eb *EmailBuilder) To(value string) *EmailBuilder {
	if !strings.Contains(value, "@") {
		eb.logger.Error().Err(ErrNoValidEmail)
		return eb
	}

	eb.email.to = value
	return eb

}

// From - used for set "from" address
func (eb *EmailBuilder) From(value string) *EmailBuilder {
	if !strings.Contains(value, "@") {
		eb.logger.Error().Err(ErrNoValidEmail)
		return eb
	}

	eb.email.from = value
	return eb
}

// Body - used for set "body" for email
func (eb *EmailBuilder) Body(value string) *EmailBuilder {
	eb.email.body = value
	return eb
}

// Subject - used for set subject of email
func (eb *EmailBuilder) Subject(value string) *EmailBuilder {
	eb.email.subject = value
	return eb
}

// GetMessage - used to get a specific product of builder
func (eb *EmailBuilder) GetMessage() Message {
	return eb.email
}

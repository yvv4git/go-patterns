package builder

import (
	"os"
	"reflect"
	"testing"

	"github.com/rs/zerolog"
)

func TestEmailBuilder_GetMessage(t *testing.T) {
	type fields struct {
		logger  zerolog.Logger
		to      string
		from    string
		subject string
		body    string
	}

	tests := []struct {
		name   string
		fields fields
		want   Message
	}{
		{
			name: "CASE-1",
			fields: fields{
				logger: zerolog.Logger{}.
					Output(os.Stdout).
					With().
					Str("project", "builder").
					Timestamp().
					Logger(),
				to:      "anonymous@gmail.com",
				from:    "ecorp@gmail.com",
				subject: "For official use",
				body:    "Hi! You are hacked!",
			},
			want: Message{
				to:      "anonymous@gmail.com",
				from:    "ecorp@gmail.com",
				subject: "For official use",
				body:    "Hi! You are hacked!",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			eb := NewEmailBuilder(tt.fields.logger)
			eb.To(tt.fields.to).
				From(tt.fields.from).
				Subject(tt.fields.subject).
				Body(tt.fields.body)

			if got := eb.GetMessage(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EmailBuilder.GetMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

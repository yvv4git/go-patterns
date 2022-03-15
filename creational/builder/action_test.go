package builder

import (
	"fmt"
	"os"
	"testing"

	"github.com/rs/zerolog"
)

func TestSendEmail(t *testing.T) {
	logger := zerolog.Logger{}.
		Output(os.Stdout).
		With().
		Str("project", "builder").
		Timestamp().
		Logger()

	type args struct {
		builder Builder
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "CASE-1",
			args: args{
				builder: func() Builder {
					b := NewEmailBuilder(logger)
					b.To("anonymous@gmail.com").
						From("ecorp@gmail.com").
						Subject("For official use").
						Body("Hi! You are hacked!")
					return &b
				}(),
			},
			want: fmt.Sprintf(
				"Processing email send ...\nTo:%s,\nSubject:%s\n%s\nRegards,\n%s\n",
				"anonymous@gmail.com",
				"For official use",
				"Hi! You are hacked!",
				"ecorp@gmail.com",
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sendEmail(tt.args.builder.GetMessage()); got != tt.want {
				t.Errorf("sendMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

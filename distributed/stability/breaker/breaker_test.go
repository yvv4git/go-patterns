package breaker

import (
	"golang.org/x/net/context"
	"testing"
	"time"
)

func TestBreaker(t *testing.T) {
	type args struct {
		failureThreshold uint          // Количество допустимых попыток
		delay            time.Duration // Интервал времени, после которого обнуляется количество попыток
	}

	testCases := []struct {
		name       string
		args       args
		ctxTimeout time.Duration
		repeat     int
	}{
		{
			name: "CASE-1",
			args: args{
				failureThreshold: 3,
				delay:            time.Second * 2,
			},
			ctxTimeout: time.Second * 10,
			repeat:     5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockActivator := NewMockActivator()
			action := Breaker(mockActivator.Activate, tc.args.failureThreshold, tc.args.delay)

			ctx := context.Background()
			ctxt, cancel := context.WithTimeout(ctx, tc.ctxTimeout)
			defer cancel()

			for i := 0; i < tc.repeat; i++ {
				_, err := action(ctxt)
				t.Log(err)
			}
		})
	}
}

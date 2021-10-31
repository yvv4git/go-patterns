package timeout

import (
	"golang.org/x/net/context"
	"testing"
	"time"
)

func TestTimeout(t *testing.T) {
	type args struct {
		mock Action
		val  int
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
				mock: NewMockActivator().Activate,
				val:  5,
			},
			ctxTimeout: time.Second * 10,
			repeat:     5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			ctxt, cancel := context.WithTimeout(ctx, tc.ctxTimeout)
			defer cancel()

			timeout := Timeout(tc.args.mock)
			for i := 0; i < tc.repeat; i++ {
				result, err := timeout(ctxt, tc.args.val)
				t.Log(result, err)
			}
		})
	}
}

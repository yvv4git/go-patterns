package debounce

import (
	"golang.org/x/net/context"
	"testing"
	"time"
)

func TestDebounceFirstTest(t *testing.T) {
	type args struct {
		mockActivator Activator     // Mock, который симулирует взаимодействие с реальным сервером
		delay         time.Duration // Интервал времени, спустя который обнуляется счетчик
	}

	testCases := []struct {
		name       string
		args       args
		ctxTimeout time.Duration // Общее ограничение времени на выполнение операции
		repeat     int           // Количество повторов, в некотором смысле это есть "дребезг"
	}{
		{
			name: "CASE-1",
			args: args{
				mockActivator: NewMockActivator(),
				delay:         time.Second * 2,
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

			action := Debounce(tc.args.mockActivator.Activate, tc.args.delay)
			for i := 0; i < tc.repeat; i++ {
				_, err := action(ctxt)
				t.Log(err)
			}
		})
	}
}

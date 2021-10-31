package retry

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"testing"
	"time"
)

func TestRetry(t *testing.T) {
	testCases := []struct {
		name          string
		retries       int           // Количество повторных попыток
		delay         time.Duration // Задержка между попытками
		ctxTimeout    time.Duration // Общее время, в течении которого допускается выполнение повторов
		expectedError error
	}{
		{
			name:          "CASE-1",
			retries:       5,
			delay:         time.Second * 3,
			ctxTimeout:    time.Second * 1,
			expectedError: context.DeadlineExceeded,
		},
		{
			name:          "CASE-2",
			retries:       5,
			delay:         time.Second * 2,
			ctxTimeout:    time.Second * 10,
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockActivator := NewMockActivator()
			retry := Retry(mockActivator.Activate, tc.retries, tc.delay)

			ctx := context.Background()
			ctxt, cancel := context.WithTimeout(ctx, tc.ctxTimeout)
			defer cancel()

			_, err := retry(ctxt)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

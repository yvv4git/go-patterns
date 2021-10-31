package throttle

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"testing"
	"time"
)

func TestThrottle(t *testing.T) {
	type args struct {
		maxAttempts    uint          // Количество допустимых повторов
		supply         uint          // Раз в некоторый период времени добавляем заданное количество попыток
		reloadInterval time.Duration // Интервал, спустя который добавляем заданное количество попыток
	}

	testCases := []struct {
		name              string
		args              args
		runCount          int            // Сколько раз запускать action
		ctxTimeout        time.Duration  // Ограничение по времени на тест, которое отменяет все попытки
		runInterval       time.Duration  // Пауза между попытками
		expectedStatistic map[string]int // Обобщенный результат по тесту.
	}{
		{
			name: "CASE-1",
			args: args{
				maxAttempts:    1,
				supply:         1,
				reloadInterval: time.Second * 3,
			},
			runCount:    5,
			ctxTimeout:  time.Second * 1,
			runInterval: time.Second * 1,
			expectedStatistic: map[string]int{
				"context deadline exceeded": 4,
				"timeout for action":        1,
			},
		},
		{
			name: "CASE-2",
			args: args{
				maxAttempts:    3,
				supply:         1,
				reloadInterval: time.Second * 3,
			},
			runCount:    5,
			ctxTimeout:  time.Second * 10,
			runInterval: time.Second * 1,
			expectedStatistic: map[string]int{
				"success":            2,
				"timeout for action": 3,
			},
		},
		{
			name: "CASE-3",
			args: args{
				maxAttempts:    2,
				supply:         1,
				reloadInterval: time.Second * 5,
			},
			runCount:    5,
			ctxTimeout:  time.Second * 10,
			runInterval: time.Second * 1,
			expectedStatistic: map[string]int{
				"throttle to many calls": 2,
				"timeout for action":     3,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockActivator := NewMockActivator()
			action := Throttle(mockActivator.Activate, tc.args.maxAttempts, tc.args.supply, tc.args.reloadInterval)

			ctx := context.Background()
			ctxt, cancel := context.WithTimeout(ctx, tc.ctxTimeout)
			defer cancel()

			statistic := make(map[string]int, 0)
			for i := 0; i < tc.runCount; i++ {
				_, err := action(ctxt)
				switch {
				case err == context.DeadlineExceeded:
					statistic["context deadline exceeded"]++
				case err == ErrToManyCalls:
					statistic["throttle to many calls"]++
				case err == ErrTimeoutAction:
					statistic["timeout for action"]++
				case err == nil:
					statistic["success"]++
				default:
					t.Log("--->", err)
					statistic["unknown result"]++
				}
				time.Sleep(tc.runInterval)
			}

			assert.Equal(t, tc.expectedStatistic, statistic)
		})
	}
}

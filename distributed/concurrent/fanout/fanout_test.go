package fanout

import (
	"sync"
	"testing"
)

func TestFanOut(t *testing.T) {
	testCases := []struct {
		name           string
		cntDstChannels int // Количество выходных каналов
	}{
		{
			name:           "CASE-1",
			cntDstChannels: 5,
		},
	}

	generator := func(n int) chan int {
		chSource := make(chan int)

		go func() {
			for i := 1; i <= n; i++ {
				chSource <- i
			}
			close(chSource)
		}()

		return chSource
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var wg sync.WaitGroup

			wg.Add(tc.cntDstChannels)
			chSource := generator(tc.cntDstChannels)

			chDestination := FanOut(chSource, tc.cntDstChannels)
			for index, singleChannel := range chDestination {
				// Можно сказать, что в отдельном потоке крутится worker
				go func(idx int, ch <-chan int) {
					defer wg.Done()

					for value := range ch {
						t.Logf("Worker[%d]:%d", idx, value)
					}
				}(index+1, singleChannel)
			}

			wg.Wait()
		})
	}
}

package fanin

import "testing"

func TestFunIn(t *testing.T) {
	type args struct {
		tasks []int
	}

	testCases := []struct {
		name string
		args args
	}{
		{
			name: "CASE-1",
			args: args{
				tasks: []int{1, 2, 3, 4, 5},
			},
		},
	}

	// Работники.
	// 1. получают задачу.
	// 2. выполняют.
	// 3. результат записывают в канал.
	worker := func(value int) chan int {
		const multiplier int = 2
		ch := make(chan int, 0)

		go func() {
			ch <- value * multiplier
			close(ch)
		}()

		return ch
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sources := make([]<-chan int, 0) // Пустой срез с каналами

			// Раздает задачи работникам
			for _, value := range tc.args.tasks {
				chSource := worker(value)
				sources = append(sources, chSource)
			}

			// Запускаем мультиплексор
			chResult := FanIn(sources...)
			for resValue := range chResult {
				t.Log(resValue)
			}
		})
	}
}

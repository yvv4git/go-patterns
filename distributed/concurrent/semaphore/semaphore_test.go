package semaphore

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestSemaphore(t *testing.T) {
	{
		// Создаем семафор на один тикет, и на ожидание освобождения светофора дается 3 секунды.
		tickets, timeout := 1, 3*time.Second
		s := New(tickets, timeout)

		// Захватываем одну задачу.
		// В данном случае допустимо захватить одновременно только 1 раз и на 3 секунды.
		require.NoError(t, s.Acquire())

		// Освобождаем семафор.
		require.NoError(t, s.Release())
	}

	{
		// Создаем семафор на один тикет, и на выполнение задачи дается 3 секунды.
		tickets, timeout := 1, 3*time.Second
		s := New(tickets, timeout)

		require.NoError(t, s.Acquire())
		require.Error(t, s.Acquire()) // Если семафор свободен, то захватываем ticket, иначе ждем 3 сек и, если не успели захватить в течении 3-х секунд ticket, возвращаем ошибку.
		require.NoError(t, s.Release())
	}

	{
		// Создаем семафор:
		// - 2 тикета, т.е. можно одновременно захватить семафор для выполнения 2-х задач.
		tickets, timeout := 2, 3*time.Second
		s := New(tickets, timeout)

		// Захватываем 2 задачи.
		require.NoError(t, s.Acquire())
		require.NoError(t, s.Acquire())

		// Освобождаем семафор.
		require.NoError(t, s.Release())
		require.NoError(t, s.Release())
	}
}

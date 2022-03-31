package lazy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLazy_Value(t *testing.T) {
	calculator := func() func() *Of[int] {
		var n = 0

		return func() *Of[int] {
			return New(func() int {
				return n + 1
			})
		}
	}

	testCases := []struct {
		name   string
		repeat int
		want   int
	}{
		{
			name:   "CASE-1",
			repeat: 2,
			want:   1,
		},
		{
			name:   "CASE-2",
			repeat: 50,
			want:   1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := 0
			for i := 0; i < tc.repeat; i++ {
				lazyBox := calculator()()
				result = lazyBox.Value()
			}

			require.Equal(t, tc.want, result)
		})
	}
}

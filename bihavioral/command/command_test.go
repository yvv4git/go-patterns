package command

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommand(t *testing.T) {
	testCases := []struct {
		name                string
		buttonStartEngine   *Button
		buttonStopEngine    *Button
		expectedResultStart Result
		expectedResultStop  Result
	}{
		{
			name: "CASE-1",
			buttonStartEngine: &Button{
				command: &EngineStartCommand{
					car: BMV{},
				},
			},
			buttonStopEngine: &Button{
				command: &EngineStopCommand{
					car: BMV{},
				},
			},
			expectedResultStart: "start",
			expectedResultStop:  "stop",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resultOfStart := tc.buttonStartEngine.Press()
			resultOfStop := tc.buttonStopEngine.Press()

			assert.Equal(t, tc.expectedResultStart, resultOfStart)
			assert.Equal(t, tc.expectedResultStop, resultOfStop)
		})
	}
}

package sharding

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShardedMapGet(t *testing.T) {
	type args struct {
		countShards     int
		indexCalculator Calculator
	}

	testCases := []struct {
		name          string
		args          args
		source        []string
		keyCheck      string
		expectedIndex int
	}{
		{
			name: "CASE-1",
			args: args{
				countShards:     7,
				indexCalculator: CalculatorMD5{},
			},
			source:        []string{"red", "orange", "yellow", "green", "blue", "indigo", "violet"},
			keyCheck:      "red",
			expectedIndex: 0,
		},
		{
			name: "CASE-2",
			args: args{
				countShards:     2,
				indexCalculator: CalculatorMD5{},
			},
			source:        []string{"red", "orange", "yellow", "green", "blue", "indigo", "violet"},
			keyCheck:      "red",
			expectedIndex: 0,
		},
		{
			name: "CASE-3",
			args: args{
				countShards:     1,
				indexCalculator: CalculatorMD5{},
			},
			source:        []string{"red", "orange", "yellow", "green", "blue", "indigo", "violet"},
			keyCheck:      "green",
			expectedIndex: 3,
		},
		{
			name: "CASE-4",
			args: args{
				countShards:     1,
				indexCalculator: CalculatorMD5{},
			},
			source:        []string{"red", "orange", "yellow", "green", "blue", "indigo", "violet"},
			keyCheck:      "blue",
			expectedIndex: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			shardMap := NewSharedMap(tc.args.countShards, tc.args.indexCalculator)
			for idx, value := range tc.source {
				shardMap.Set(value, idx)
			}

			actualValue := shardMap.Get(tc.keyCheck)
			assert.Equal(t, tc.expectedIndex, actualValue)
		})
	}
}

func TestSharedMapKeys(t *testing.T) {
	type args struct {
		countShards     int
		indexCalculator Calculator
	}

	testCases := []struct {
		name              string
		source            []string
		args              args
		expectedKeysCount int
	}{
		{
			name:   "CASE-1",
			source: []string{"red", "orange", "yellow", "green", "blue", "indigo", "violet"},
			args: args{
				countShards:     5,
				indexCalculator: CalculatorMD5{},
			},
			expectedKeysCount: 7,
		},
		{
			name:   "CASE-2",
			source: []string{"red", "orange", "yellow"},
			args: args{
				countShards:     5,
				indexCalculator: CalculatorMD5{},
			},
			expectedKeysCount: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			shardMap := NewSharedMap(tc.args.countShards, tc.args.indexCalculator)
			for idx, value := range tc.source {
				shardMap.Set(value, idx)
			}

			keys := shardMap.Keys()
			assert.Equal(t, tc.expectedKeysCount, len(keys))
		})
	}
}

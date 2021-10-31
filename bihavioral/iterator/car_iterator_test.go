package iterator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCarIterator(t *testing.T) {
	testCases := []struct {
		name             string
		carList          CarCollection
		expectedCarNames []string
	}{
		{
			name: "CASE-1",
			carList: CarCollection{
				&Car{
					Name:  "Vesta",
					Mark:  "Lada",
					Age:   2021,
					Price: 100000,
				},
				&Car{
					Name:  "Logan",
					Mark:  "Renault",
					Age:   2021,
					Price: 1300000,
				},
			},
			expectedCarNames: []string{
				"Vesta",
				"Logan",
			},
		},
	}

	clientCode := func(iterator Iterator) []string {
		var resultCarNames []string
		for iterator.Next() {
			car := iterator.Get().(Car)
			resultCarNames = append(resultCarNames, car.Name)
		}
		return resultCarNames
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			carIterator := NewCarIterator(tc.carList)
			resultCarNames := clientCode(carIterator)
			assert.Equal(t, tc.expectedCarNames, resultCarNames)
		})
	}
}

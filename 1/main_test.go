package main

import (
	"testing"
)

func TestGetSum(t *testing.T) {
	testCases := []struct {
		name        string
		inputValues []string
		expectedSum int
	}{
		{
			name:        "Valid values",
			inputValues: []string{"two1nine", "eightwothree", "abcone2threexyz", "xtwone3four", "4nineeightseven2", "zoneight234", "7pqrstsixteen"},
			expectedSum: 281,
		},
		{
			name:        "Valid values",
			inputValues: []string{"flhmdp6eighteightmcxcvffive"},
			expectedSum: 65,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			sum := GetSum(testCase.inputValues)

			if sum != testCase.expectedSum {
				t.Errorf("Expected sum %d, but got %d", testCase.expectedSum, sum)
			}
		})
	}
}

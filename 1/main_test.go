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
			name:        "Handle digits",
			inputValues: []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"},
			expectedSum: 142,
		},
		{
			name:        "Handle digit and number in words ",
			inputValues: []string{"two1nine", "eightwothree", "abcone2threexyz", "xtwone3four", "4nineeightseven2", "zoneight234", "7pqrstsixteen"},
			expectedSum: 281,
		},
		{
			name:        "Handle numeronym (eg: twone for 2 and 1, threeight for 3 and 8)",
			inputValues: []string{"flhmtwoonexc", "threeight"},
			expectedSum: 59,
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

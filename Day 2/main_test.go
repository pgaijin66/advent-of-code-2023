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
			name: "Sum of game ID of all valid games",
			inputValues: []string{
				"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
				"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
				"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
				"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
				"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			},
			expectedSum: 8,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			sum := GetSumOfValidGameID(testCase.inputValues)

			if sum != testCase.expectedSum {
				t.Errorf("Expected sum %d, but got %d", testCase.expectedSum, sum)
			}
		})
	}
}

func TestGetMin(t *testing.T) {
	testCases := []struct {
		name        string
		inputValues []string
		expectedSum int
	}{
		{
			name: "Sum of power of minimum set of cubes required to have a valid game",
			inputValues: []string{
				"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
				"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
				"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
				"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
				"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			},
			expectedSum: 2286,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			sum := GetSumOfPower(testCase.inputValues)

			if sum != testCase.expectedSum {
				t.Errorf("Expected sum %d, but got %d", testCase.expectedSum, sum)
			}
		})
	}
}

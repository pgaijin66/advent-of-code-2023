package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

// Number represents a custom integer type.
type Number int

var (
	inputValues       []string
	calibrationValues []int
)

var (
	// NumberNames maps custom Number type to their string representations.
	NumberNames = map[Number]string{
		0: "zero",
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
		7: "seven",
		8: "eight",
		9: "nine",
	}
	// ReverseNumberNames maps lowercase string representations to custom Number type for reverse lookup.
	ReverseNumberNames map[string]Number
)

func init() {
	// Read input from file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %s", err)
	}
	defer file.Close()

	// Scan individual files
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputValues = append(inputValues, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	// Create reverse number name map for easier lookup
	initializeReverseNumberNames()
}

func main() {
	fmt.Println(CalculateSum(inputValues))
}

func CalculateSum(values []string) int {
	var sum int

	for _, value := range values {
		value, err := getCalibrationValues(value)
		if err != nil {
			log.Fatalf("could not get calibration value: %s", err)
		}
		calibrationValues = append(calibrationValues, value)
		sum += value

	}

	return sum
}

func getCalibrationValues(input string) (int, error) {
	var firstDigit, lastDigit int
	// var foundFirstDigit bool
	var CalibrationNumbers []int

	for _, char := range input {
		switch {
		case unicode.IsDigit(char):
			digitValue := int(char - '0')
			CalibrationNumbers = append(CalibrationNumbers, digitValue)

		case isCharStartForAnyNumber(char):
			a := findNumberFromWords(string(input))
			if a != 0 {
				CalibrationNumbers = append(CalibrationNumbers, a)
			}
		}

		indexOfChar := strings.Index(input, string(char))
		input = input[indexOfChar+1:]
	}

	tmp := removeZeros(CalibrationNumbers)
	firstDigit = tmp[0]
	lastDigit = tmp[len(tmp)-1]

	return firstDigit*10 + lastDigit, nil
}

func initializeReverseNumberNames() {
	ReverseNumberNames = make(map[string]Number)
	for num, name := range NumberNames {
		ReverseNumberNames[strings.ToLower(name)] = num
	}
}

func findNumberFromWords(input string) int {
	var currentNumberInWord string
	var number int

	for _, char := range input {
		if unicode.IsLetter(char) {
			currentNumberInWord += string(char)
			if isValidWord(currentNumberInWord) {
				number = int(ReverseNumberNames[strings.ToLower(currentNumberInWord)])
				// fmt.Println("Found a valid word: ", number)
				return number
			}
		}
	}

	return number
}

func isCharStartForAnyNumber(char rune) bool {
	lowerChar := unicode.ToLower(char)
	for _, word := range NumberNames {
		if strings.HasPrefix(strings.ToLower(word), string(lowerChar)) {
			return true
		}
	}
	return false
}

func isValidWord(word string) bool {
	_, ok := ReverseNumberNames[strings.ToLower(word)]
	return ok
}

func removeZeros(input []int) []int {
	var result []int
	for _, num := range input {
		if num != 0 {
			result = append(result, num)
		}
	}
	return result
}

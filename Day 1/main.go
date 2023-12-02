package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

type Number int

var (
	inputValues []string
)

var (
	WordToDigit = map[string]Number{
		"nine":  9,
		"eight": 8,
		"seven": 7,
		"six":   6,
		"five":  5,
		"four":  4,
		"three": 3,
		"two":   2,
		"one":   1,
		"zero":  0,
	}
)

func init() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputValues = append(inputValues, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err)
	}
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
		sum += value

	}

	return sum
}

func getCalibrationValues(input string) (int, error) {
	var firstDigit, lastDigit int
	var calibrationValue []int

	for _, char := range input {
		switch {
		case unicode.IsDigit(char):
			digitValue := int(char - '0')
			calibrationValue = append(calibrationValue, digitValue)

		case unicode.IsLetter(char) && isCharStartForAnyNumber(char):
			num := findNumberFromWords(string(input))
			if num != 0 {
				calibrationValue = append(calibrationValue, num)
			}

		}

		indexOfChar := strings.Index(input, string(char))
		input = input[indexOfChar+1:]
	}

	firstDigit = calibrationValue[0]
	lastDigit = calibrationValue[len(calibrationValue)-1]

	return firstDigit*10 + lastDigit, nil
}

func findNumberFromWords(input string) int {
	var currentNumberInWord string
	var number int

	for _, char := range input {
		currentNumberInWord += string(char)
		if isValidWord(currentNumberInWord) {
			number = int(WordToDigit[strings.ToLower(currentNumberInWord)])
			return number
		}
	}

	return number
}

func isCharStartForAnyNumber(char rune) bool {
	lowerChar := unicode.ToLower(char)
	for word := range WordToDigit {
		if strings.HasPrefix(strings.ToLower(word), string(lowerChar)) {
			return true
		}
	}
	return false
}

func isValidWord(word string) bool {
	_, ok := WordToDigit[strings.ToLower(word)]
	return ok
}

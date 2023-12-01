package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

var inputValues []string

var calibrationValues []int

type Number int

var NumberNames = map[Number]string{
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

var ReverseNumberNames map[string]Number

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

	ReverseNumberNames = make(map[string]Number)
	for num, name := range NumberNames {
		ReverseNumberNames[strings.ToLower(name)] = num
	}
}

func isValidWord(word string) bool {
	_, ok := ReverseNumberNames[strings.ToLower(word)]
	return ok
}

func IsCharStartForAnyNumber(char rune) bool {
	lowerChar := unicode.ToLower(char)
	for _, word := range NumberNames {
		if strings.HasPrefix(strings.ToLower(word), string(lowerChar)) {
			return true
		}
	}
	return false
}

func GetCalibrationValues(input string) (int, error) {
	var firstDigit, lastDigit int
	var foundFirstDigit bool
	var numbers []int
	var remainingString string

	for _, char := range input {
		fmt.Println("Checking char: ", string(char))
		switch {
		case unicode.IsDigit(char):
			digitValue := int(char - '0')
			if !foundFirstDigit {
				firstDigit = digitValue
				foundFirstDigit = true
			}
			lastDigit = digitValue
			numbers = append(numbers, digitValue)

		case IsCharStartForAnyNumber(char):
			fmt.Println("Found: ", string(char))
			fmt.Println("Checking Substring: ", input)
			indexOfChar := strings.Index(input, string(char))
			remainingString = input[indexOfChar:]
			a := FindNumberFromWords(string(remainingString))
			if a != 0 {
				numbers = append(numbers, a)
			}
			input = remainingString

		default:
			indexOfChar := strings.Index(input, string(char))
			remainingString = input[indexOfChar:]
			input = remainingString
			fmt.Println("new string", input)
		}

	}

	newArray := removeZeros(numbers)
	firstDigit = newArray[0]
	lastDigit = newArray[len(newArray)-1]
	result := firstDigit*10 + lastDigit

	return result, nil
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

func FindNumberFromWords(input string) int {
	var currentNumberInWord string
	var number int

	for _, char := range input {
		if unicode.IsLetter(char) {
			currentNumberInWord += string(char)
			if isValidWord(currentNumberInWord) {
				number = int(ReverseNumberNames[strings.ToLower(currentNumberInWord)])
				return number
			}
		}
	}

	return number
}

func GetSum(values []string) int {
	var sum int

	for _, value := range values {
		value, err := GetCalibrationValues(value)
		if err != nil {
			log.Fatalf("could not get calibration value: %s", err)
		}
		calibrationValues = append(calibrationValues, value)
		sum += value

	}
	fmt.Println(calibrationValues)

	return sum
}

func main() {
	fmt.Println(GetSum(inputValues))
}

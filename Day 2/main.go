package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	inputValues []string
)

var (
	TotalCubeCount = map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
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
}

func GetSumOfValidGameID(inputValues []string) int {
	var sum int

	for _, value := range inputValues {
		sets := strings.Split(value, ": ")
		gameID, err := strconv.Atoi(strings.Split(sets[0], " ")[1])
		if err != nil {
			log.Fatal(err)
		}

		diceSets := strings.Split(sets[1], "; ")

		if isValidGame(diceSets) {
			sum += gameID
		}
	}

	return sum
}

func GetSumOfPower(inputValues []string) int {
	var minResult int

	for _, value := range inputValues {
		sets := strings.Split(value, ": ")
		diceSets := strings.Split(sets[1], "; ")

		minCube := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, dieSet := range diceSets {
			dice := strings.Split(dieSet, ", ")
			for _, die := range dice {
				number, color, err := parseDice(die)
				if err != nil {
					log.Fatal(err)
				}

				if number > minCube[color] {
					minCube[color] = number
				}
			}
		}

		power := minCube["red"] * minCube["green"] * minCube["blue"]
		minResult += power
	}

	return minResult
}

func main() {
	SumOfValidGameID := GetSumOfValidGameID(inputValues)
	SumOfPower := GetSumOfPower(inputValues)
	fmt.Println(SumOfValidGameID, SumOfPower)

}

func isValidGame(diceSets []string) bool {
	for _, dieSet := range diceSets {
		dice := strings.Split(dieSet, ", ")
		for _, die := range dice {
			number, color, err := parseDice(die)
			if err != nil {
				log.Fatal(err)
			}

			if number > TotalCubeCount[color] {
				return false
			}
		}
	}
	return true
}

func parseDice(die string) (int, string, error) {
	parts := strings.Split(die, " ")
	if len(parts) != 2 {
		return 0, "", fmt.Errorf("invalid dice format: %s", die)
	}

	number, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, "", fmt.Errorf("could not convert dice string to int: %s", err)
	}

	return number, parts[1], nil
}

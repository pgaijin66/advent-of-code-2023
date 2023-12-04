package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	WinningNumbers []int
	NumberOnHand   []int
	Points         int
}

type Cards []Card

func main() {
	var totalPoints int = 0
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("could not open file")
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var card Card
		// Initial card point will be 0
		card.Points = 0
		a := strings.Split(scanner.Text(), ": ")[1]

		wn := strings.Split(strings.TrimSpace(strings.Split(a, "| ")[0]), " ")
		nh := strings.Split(strings.TrimSpace(strings.Split(a, "| ")[1]), " ")

		card.WinningNumbers, err = convertStringsToNumbers(wn)
		if err != nil {
			log.Fatal(err)
		}
		card.NumberOnHand, _ = convertStringsToNumbers(nh)

		for _, num := range card.NumberOnHand {
			if card.Points != 0 && isNumberExists(card.WinningNumbers, num) {
				card.Points *= 2
			} else if card.Points == 0 && isNumberExists(card.WinningNumbers, num) {
				card.Points += 1
			}
		}
		totalPoints = totalPoints + card.Points

	}

	fmt.Println("[Part 1] Total card points: ", totalPoints)
}

func isNumberExists(arr []int, numToCheck int) bool {
	// fmt.Println("Checking if ", numToCheck, " exists in ", arr)
	for _, num := range arr {
		if num == numToCheck {
			return true
		}
	}
	return false
}

func convertStringsToNumbers(strArray []string) ([]int, error) {
	var intArray []int

	for _, str := range strArray {
		if str == "" {
			continue
		} else {
			num, err := strconv.Atoi(strings.TrimSpace(str))
			if err != nil {
				return nil, err
			}
			intArray = append(intArray, num)
		}

	}

	return intArray, nil
}

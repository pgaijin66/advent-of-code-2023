package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := ReadFile("input.txt")
	fmt.Println("first part:")
	schematic := NewSchematic(input)
	partNumbersSum := SumPartNumbers(schematic.GetSchematicParts().Numbers)
	fmt.Println(partNumbersSum)

	fmt.Println("second part:")
	gearCandidates := schematic.GetSchematicParts().GearCandidates
	fmt.Printf("%d\n", SumAllGearRatios(gearCandidates))
}

type SchematicNumber struct {
	Value           int
	AdjacentSymbols map[rune]bool
}

type SchematicDimensions struct {
	Width  int
	Length int
}

type Schematic struct {
	Contents   string
	Dimensions SchematicDimensions
}

func NewSchematic(input string) Schematic {
	schematic := Schematic{Contents: input}
	for i, c := range input {
		if c == rune('\n') {
			schematic.Dimensions.Width = i
			schematic.Dimensions.Length = len(input) / schematic.Dimensions.Width
			break
		}
	}
	return schematic
}

var ErrOutOfBounds = fmt.Errorf("symbol out of bounds")

func (s Schematic) GetSymbol(x, y int) (rune, error) {
	if x >= s.Dimensions.Width || y >= s.Dimensions.Length || x < 0 || y < 0 {
		return 0, ErrOutOfBounds
	}
	symbolIndex := x + s.Dimensions.Width*y + y
	return rune(s.Contents[symbolIndex]), nil
}

func NewSchematicNumber(value int) SchematicNumber {
	return SchematicNumber{Value: value, AdjacentSymbols: make(map[rune]bool)}
}

type Coordinates struct {
	x, y int
}

type GearCandidate []int

type GearCandidates map[Coordinates]GearCandidate

func (g GearCandidate) IsGear() bool {
	return len(g) == 2
}

func (g GearCandidate) GetRatio() int {
	ratio := 1
	for _, value := range g {
		ratio *= value
	}
	return ratio
}

func (g GearCandidates) SumAllGearRatios() int {
	sumRatios := 0
	for _, gearCandidate := range g {
		if gearCandidate.IsGear() {
			sumRatios += gearCandidate.GetRatio()
		}
	}
	return sumRatios
}

func (s Schematic) GetSchematicParts() struct {
	Numbers        []SchematicNumber
	GearCandidates GearCandidates
} {
	numbers := []SchematicNumber{}
	gearCandidates := make(GearCandidates)

	x, y := 0, 0

	incrementCoordinates := func(c rune) {
		if c == rune('\n') {
			x, y = 0, y+1
		} else {
			x++
		}
	}

	currentValue := 0
	numberLength := 0

	newSchematicNumberWithAdjacentSymbols := func() SchematicNumber {
		number := NewSchematicNumber(currentValue)
		addSymbolIfPossible := func(cx, cy int) {
			if symbol, err := s.GetSymbol(cx, cy); err == nil {
				number.AdjacentSymbols[symbol] = true
				if symbol == rune('*') {
					gearCandidates[Coordinates{x: cx, y: cy}] = append(gearCandidates[Coordinates{x: cx, y: cy}], currentValue)
				}
			}
		}
		addSymbolIfPossible(x, y)
		addSymbolIfPossible(x-numberLength-1, y)
		for i := x - numberLength - 1; i <= x; i++ {
			addSymbolIfPossible(i, y-1)
			addSymbolIfPossible(i, y+1)
		}
		return number
	}

	for _, char := range s.Contents {
		if digit, err := strconv.Atoi(string(char)); err == nil {
			currentValue = currentValue*10 + digit
			numberLength++
		} else if currentValue > 0 {
			currentNumber := newSchematicNumberWithAdjacentSymbols()
			numbers = append(numbers, currentNumber)
			currentValue = 0
			numberLength = 0
		}
		incrementCoordinates(char)
	}

	return struct {
		Numbers        []SchematicNumber
		GearCandidates GearCandidates
	}{Numbers: numbers, GearCandidates: gearCandidates}
}

func (n SchematicNumber) IsPartNumber() bool {
	for symbol := range n.AdjacentSymbols {
		if symbol != '.' {
			return true
		}
	}
	return false
}

func SumPartNumbers(schematicNumbers []SchematicNumber) int {
	sum := 0
	for _, schematicNumber := range schematicNumbers {
		if schematicNumber.IsPartNumber() {
			sum += schematicNumber.Value
		}
	}
	return sum
}

func SumAllGearRatios(gearCandidates GearCandidates) int {
	sumRatios := 0
	for _, gearCandidate := range gearCandidates {
		if gearCandidate.IsGear() {
			sumRatios += gearCandidate.GetRatio()
		}
	}
	return sumRatios
}

func ReadFile(fileName string) string {
	content, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(content))
}

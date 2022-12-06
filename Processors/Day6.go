package Processors

import (
	"AOC_2022/FileReader"
	"fmt"
	"strings"
)

type Day6 struct {
	Processor
	InputFilePath string
}

func (Day6 Day6) FirstTaskResult() {
	fmt.Println("First task:")
	inputString := FileReader.GetString(Day6.InputFilePath)

	characters := strings.Split(inputString, "")

	result := getFirstUniqueOccurrence(characters, 4)
	fmt.Println(result)
}

func getFirstUniqueOccurrence(characters []string, length int) int {
	var result int
	for i := 0; i < len(characters)-length; i++ {
		var chars []string
		for y := 0; y < length; y++ {
			chars = append(chars, characters[y+i])
		}

		if !containsDuplicate(chars) {
			result = i + length
			break
		}
	}

	return result
}

func containsDuplicate(stringArray []string) bool {
	valueMap := make(map[string]string)
	for _, value := range stringArray {
		if valueMap[value] != "" {
			return true
		} else {
			valueMap[value] = value
		}
	}

	return false
}

func (Day6 Day6) SecondTaskResult() {
	fmt.Println("First task:")
	inputString := FileReader.GetString(Day6.InputFilePath)

	characters := strings.Split(inputString, "")

	result := getFirstUniqueOccurrence(characters, 14)
	fmt.Println(result)
}

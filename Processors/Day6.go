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

	var result int
	for i := 0; i < len(characters)-4; i++ {
		chars := []string{characters[i], characters[i+1], characters[i+2], characters[i+3]}

		if !containsDuplicate(chars) {
			result = i + 4
			break
		}
	}

	fmt.Println(result)
}

func containsDuplicate(stringArray []string) bool {
	valueMap := make(map[string]string)
	for _, value := range stringArray {
		if valueMap[value] != "" {
			fmt.Println("duplicate:")
			fmt.Println(stringArray)
			return true
		} else {
			valueMap[value] = value
		}
	}

	fmt.Println("No duplicate:")
	fmt.Println(stringArray)
	return false
}

func (Day6 Day6) SecondTaskResult() {
	fmt.Println("Second task:")
}

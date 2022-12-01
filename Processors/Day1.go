package Processors

import (
	fileReader "AOC_2022/FileReader"
	"fmt"
)

type Day1 struct {
	InputFilePath string
	processor     Processor
}

func (Processor Day1) FirstTaskResult() {
	fmt.Println("Processing day one task one!")
	inputMap := fileReader.GetIntArray(Processor.InputFilePath)

	maxAmount := 0
	currentSum := 0
	for _, value := range inputMap {
		if value == 0 {
			if maxAmount < currentSum {
				maxAmount = currentSum
			}
			currentSum = 0
		}

		currentSum += value
	}

	fmt.Println(maxAmount)
}

func (Processor Day1) SecondTaskResult() {
	fmt.Println("Processing day one second task!")

	fmt.Println("Processing day one task one!")
	inputMap := fileReader.GetIntArray(Processor.InputFilePath)

	var maxAmounts = make([]int, 3)

	currentSum := 0
	for _, value := range inputMap {
		if value == 0 {
			maxAmounts = getAdjustedMaxAmounts(maxAmounts, currentSum)
			currentSum = 0
		}

		currentSum += value
	}

	sum := 0
	for _, amount := range maxAmounts {
		sum += amount
	}

	fmt.Println(sum)
}

func getAdjustedMaxAmounts(maxAmounts []int, newValue int) []int {
	for i, maxAmount := range maxAmounts {
		if maxAmount < newValue {
			maxAmounts[i] = newValue
			newValue = maxAmount
		}
	}

	return maxAmounts
}

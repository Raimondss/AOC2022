package Processors

import (
	"AOC_2022/FileReader"
	"fmt"
	"strconv"
	"strings"
)

//[R] [J] [W]
//[R] [N]     [T] [T] [C]
//[R]         [P] [G]     [J] [P] [T]
//[Q]     [C] [M] [V]     [F] [F] [H]
//[G] [P] [M] [S] [Z]     [Z] [C] [Q]
//[P] [C] [P] [Q] [J] [J] [P] [H] [Z]
//[C] [T] [H] [T] [H] [P] [G] [L] [V]
//[F] [W] [B] [L] [P] [D] [L] [N] [G]
//1   2   3   4   5   6   7   8   9

type Day5 struct {
	Processor
	InputFilePath string
}

func (Day5 Day5) FirstTaskResult() {
	fmt.Println("firs task")

	commands := FileReader.GetStringArray(Day5.InputFilePath)

	stackMap := createInitialStacks()

	for _, command := range commands {
		commandArray := strings.Split(command, " ")

		itemCount, _ := strconv.Atoi(commandArray[1])
		fromStackIndex, _ := strconv.Atoi(commandArray[3])
		stackToAddIndex, _ := strconv.Atoi(commandArray[5])

		stackToTakeFrom := stackMap[fromStackIndex]
		stackToAddTo := stackMap[stackToAddIndex]

		takenElements := stackToTakeFrom.takeElementsFromStack(itemCount)
		stackToAddTo.addElementsToStack(takenElements)

		stackMap[fromStackIndex] = stackToTakeFrom
		stackMap[stackToAddIndex] = stackToAddTo
	}

	fmt.Println("Result stack:")
	fmt.Println(stackMap)

	resultString := ""
	for i := 0; i < len(stackMap); i++ {
		resultString = resultString + stackMap[i+1].elements[0]
	}

	println(resultString)

}

func (Day5 Day5) SecondTaskResult() {
	fmt.Println("firs task")

	commands := FileReader.GetStringArray(Day5.InputFilePath)

	stackMap := createInitialStacks()

	for _, command := range commands {
		commandArray := strings.Split(command, " ")

		itemCount, _ := strconv.Atoi(commandArray[1])
		fromStackIndex, _ := strconv.Atoi(commandArray[3])
		stackToAddIndex, _ := strconv.Atoi(commandArray[5])

		stackToTakeFrom := stackMap[fromStackIndex]
		stackToAddTo := stackMap[stackToAddIndex]

		takenElements := stackToTakeFrom.takeElementsFromStackInSameOrder(itemCount)
		stackToAddTo.addElementsToStack(takenElements)

		stackMap[fromStackIndex] = stackToTakeFrom
		stackMap[stackToAddIndex] = stackToAddTo
	}

	fmt.Println("Result stack:")
	fmt.Println(stackMap)

	resultString := ""
	for i := 0; i < len(stackMap); i++ {
		resultString = resultString + stackMap[i+1].elements[0]
	}

	println(resultString)

}

func createInitialStacks() map[int]Stack {
	stackMap := make(map[int]Stack)
	stackMap[1] = Stack{elements: []string{"R", "Q", "G", "P", "C", "F"}}
	stackMap[2] = Stack{elements: []string{"P", "C", "T", "W"}}
	stackMap[3] = Stack{elements: []string{"C", "M", "P", "H", "B"}}
	stackMap[4] = Stack{elements: []string{"R", "P", "M", "S", "Q", "T", "L"}}
	stackMap[5] = Stack{elements: []string{"N", "G", "V", "Z", "J", "H", "P"}}
	stackMap[6] = Stack{elements: []string{"J", "P", "D"}}
	stackMap[7] = Stack{elements: []string{"R", "T", "J", "F", "Z", "P", "G", "L"}}
	stackMap[8] = Stack{elements: []string{"J", "T", "P", "F", "C", "H", "L", "N"}}
	stackMap[9] = Stack{elements: []string{"W", "C", "T", "H", "Q", "Z", "V", "G"}}

	return stackMap
}

type Stack struct {
	elements []string
}

func (Stack *Stack) takeElementsFromStack(count int) []string {
	var takenElements []string

	var newElements []string

	for i, element := range Stack.elements {
		if i+1 <= count {
			takenElements = append(takenElements, element)
		} else {
			newElements = append(newElements, element)
		}
	}

	Stack.elements = newElements

	return reverseTakenElements(takenElements)
}

func (Stack *Stack) takeElementsFromStackInSameOrder(count int) []string {
	var takenElements []string

	var newElements []string

	for i, element := range Stack.elements {
		if i+1 <= count {
			takenElements = append(takenElements, element)
		} else {
			newElements = append(newElements, element)
		}
	}

	Stack.elements = newElements

	return takenElements
}

func (Stack *Stack) addElementsToStack(elementsToAdd []string) {

	var newStack []string

	for _, element := range elementsToAdd {
		newStack = append(newStack, element)
	}

	for _, element := range Stack.elements {
		newStack = append(newStack, element)
	}

	Stack.elements = newStack
}

func reverseTakenElements(takenElements []string) []string {
	if len(takenElements) == 0 {
		return takenElements
	}
	return append(reverseTakenElements(takenElements[1:]), takenElements[0])
}

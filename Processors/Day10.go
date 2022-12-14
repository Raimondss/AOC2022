package Processors

import (
	"AOC_2022/FileReader"
	"fmt"
	"strconv"
	"strings"
)

type Day10 struct {
	Processor
	InputFilePath string
}

func getInput(day10 Day10) []string {
	return FileReader.GetStringArray(day10.InputFilePath)
}

func (Day10 Day10) FirstTaskResult() {
	fmt.Println("First:")
	inputStrings := FileReader.GetStringArray(Day10.InputFilePath)
	var x = 1
	var cycle = 1

	totalSignalStrenth := 0
	for _, inputString := range inputStrings {
		printCycle(x, cycle)
		fmt.Println(inputString)

		command := strings.Split(inputString, " ")
		cyclesToRun := 0
		valueToAdd := 0

		if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
			signalStrength := cycle * x
			fmt.Println("Signal strenght:" + strconv.Itoa(signalStrength))
			totalSignalStrenth += signalStrength
			fmt.Println("SINGAL STRENGHT UPDATE:" + strconv.Itoa(cycle) + " " + strconv.Itoa(totalSignalStrenth))
		}

		if command[0] == "addx" {
			valueToAdd, _ = strconv.Atoi(command[1])
			cyclesToRun = 1
		}

		for i := 0; i < cyclesToRun; i++ {
			cycle++

			if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
				signalStrength := cycle * x

				totalSignalStrenth += signalStrength
				fmt.Println("Signal strenght:" + strconv.Itoa(signalStrength))
				fmt.Println("SINGAL STRENGHT UPDATE:" + strconv.Itoa(cycle) + " " + strconv.Itoa(totalSignalStrenth))
			}

			printCycle(x, cycle)
		}

		cycle++

		println("Adding value:" + strconv.Itoa(valueToAdd))
		x += valueToAdd
	}

	fmt.Println("TOTAL:" + strconv.Itoa(totalSignalStrenth))
}

func printCycle(x int, cycle int) {
	fmt.Println("CYCLE " + strconv.Itoa(cycle) + " X:" + strconv.Itoa(x))
}

func (Day10 Day10) SecondTaskResult() {
	fmt.Println("Second:")

	fmt.Println("First:")
	inputStrings := FileReader.GetStringArray(Day10.InputFilePath)
	var x = 1
	var cycle = 1

	totalSignalStrength := 0

	var pixels []string

	for _, inputString := range inputStrings {
		printCycle(x, cycle)
		fmt.Println(inputString)
		pixels = append(pixels, getPixel(x, cycle))

		//pixels = append(pixels, getPixel(x, cycle))

		command := strings.Split(inputString, " ")
		cyclesToRun := 0
		valueToAdd := 0

		if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
			signalStrength := cycle * x
			fmt.Println("Signal strenght:" + strconv.Itoa(signalStrength))
			totalSignalStrength += signalStrength
			fmt.Println("SINGAL STRENGHT UPDATE:" + strconv.Itoa(cycle) + " " + strconv.Itoa(totalSignalStrength))
		}

		if command[0] == "addx" {
			valueToAdd, _ = strconv.Atoi(command[1])
			cyclesToRun = 1
		}

		for i := 0; i < cyclesToRun; i++ {
			cycle++
			pixels = append(pixels, getPixel(x, cycle))

			if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
				signalStrength := cycle * x

				totalSignalStrength += signalStrength
				fmt.Println("Signal strenght:" + strconv.Itoa(signalStrength))
				fmt.Println("SINGAL STRENGHT UPDATE:" + strconv.Itoa(cycle) + " " + strconv.Itoa(totalSignalStrength))
			}

			printCycle(x, cycle)
		}

		cycle++

		println("Adding value:" + strconv.Itoa(valueToAdd))
		x += valueToAdd
	}

	fmt.Println(pixels)
	fmt.Println("TOTAL:" + strconv.Itoa(totalSignalStrength))

	for i, pixel := range pixels {
		if i%40 == 0 {
			println("")
		}

		print(pixel)
	}
}

func getPixel(x int, cycle int) string {
	if cycle > 40 {
		cycle = cycle - 40
	}

	if cycle >= 40 {
		cycle = cycle - 40
	}
	if cycle >= 40 {
		cycle = cycle - 40
	}
	if cycle >= 40 {
		cycle = cycle - 40
	}
	if cycle >= 40 {
		cycle = cycle - 40
	}
	if cycle >= 40 {
		cycle = cycle - 40
	}

	if cycle == x || cycle == x+2 || cycle == x+1 {
		println("DRAW # at:" + strconv.Itoa(cycle))
		return "#"
	}
	println("DRAW . at:" + strconv.Itoa(cycle))
	return "."
}

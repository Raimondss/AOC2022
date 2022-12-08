package main

import (
	"AOC_2022/Processors"
	"strconv"
)

const FilesFolder = "InputFiles"

// DAY TODO User input
const DAY = 8

func main() {
	getDayProcessor(DAY).FirstTaskResult()
	getDayProcessor(DAY).SecondTaskResult()
}

func getDayProcessor(dayNumber int) Processors.Processor {
	return getAllProcessors()[dayNumber]
}

func getAllProcessors() map[int]Processors.Processor {
	return map[int]Processors.Processor{
		1: Processors.Day1{InputFilePath: getDayInputFile(1)},
		2: Processors.Day2{InputFilePath: getDayInputFile(2)},
		3: Processors.Day3{InputFilePath: getDayInputFile(3)},
		4: Processors.Day4{InputFilePath: getDayInputFile(4)},
		5: Processors.Day5{InputFilePath: getDayInputFile(5)},
		6: Processors.Day6{InputFilePath: getDayInputFile(6)},
		7: Processors.Day7{InputFilePath: getDayInputFile(7)},
		8: Processors.Day8{InputFilePath: getDayInputFile(8)},
	}
}

func getDayInputFile(day int) string {
	dayNumber := strconv.Itoa(day)

	return FilesFolder + "/" + dayNumber
}

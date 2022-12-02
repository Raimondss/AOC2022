package main

import (
	"AOC_2022/Processors"
	"strconv"
)

const FilesFolder = "InputFiles"

// DAY TODO User input
const DAY = 2

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
	}
}

func getDayInputFile(day int) string {
	dayNumber := strconv.Itoa(day)

	return FilesFolder + "/" + dayNumber
}

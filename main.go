package main

import (
	"AOC_2022/Processors"
	"strconv"
)

const FilesFolder = "InputFiles"

func main() {
	getDayProcessor(1).FirstTaskResult()
	getDayProcessor(1).SecondTaskResult()
}

func getDayProcessor(dayNumber int) Processors.Processor {
	return getAllProcessors()[dayNumber]
}

func getAllProcessors() map[int]Processors.Processor {
	return map[int]Processors.Processor{
		1: Processors.Day1{InputFilePath: getDayInputFile(1)},
	}
}

func getDayInputFile(day int) string {
	dayNumber := strconv.Itoa(day)

	return FilesFolder + "/" + dayNumber
}

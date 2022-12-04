package Processors

import (
	"AOC_2022/FileReader"
	"fmt"
	"strconv"
	"strings"
)

type Day4 struct {
	InputFilePath string
	Processor
}

type Elf struct {
	sectionStart int
	sectionEnd   int
}

type ElfPair struct {
	firstElf  Elf
	secondElf Elf
}

func (ElfPair ElfPair) rangeContains() bool {
	return areRangesOverlapping(ElfPair.firstElf.sectionStart, ElfPair.firstElf.sectionEnd, ElfPair.secondElf.sectionStart, ElfPair.secondElf.sectionEnd)
}

func (ElfPair ElfPair) rangeFullyContains() bool {
	if rangesFullyContain(ElfPair.firstElf.sectionStart, ElfPair.firstElf.sectionEnd, ElfPair.secondElf.sectionStart, ElfPair.secondElf.sectionEnd) {
		return true
	}

	return false
}

func rangesFullyContain(x1 int, x2 int, y1 int, y2 int) bool {
	if x1 >= y1 && x2 <= y2 {
		return true
	}

	if y1 >= x1 && y2 <= x2 {
		return true
	}

	return false
}

func areRangesOverlapping(x1 int, x2 int, y1 int, y2 int) bool {
	if rangeContains(x1, x2, y1) {
		return true
	}

	if rangeContains(x1, x2, y2) {
		return true
	}

	if rangeContains(y1, y2, x1) {
		return true
	}

	if rangeContains(y1, y2, x2) {
		return true
	}

	return false
}

func rangeContains(rangeStart int, rangeEnd int, value int) bool {
	if value >= rangeStart && value <= rangeEnd {
		return true
	}

	return false
}

func createElfPairsFromInput(Day4 Day4) []ElfPair {
	pairStrings := FileReader.GetStringArray(Day4.InputFilePath)

	elfPairs := make([]ElfPair, len(pairStrings))
	for i, pairString := range pairStrings {
		elfRangeStrings := strings.Split(pairString, ",")

		firstElf := getElfFromRangeString(elfRangeStrings[0])
		secondElf := getElfFromRangeString(elfRangeStrings[1])

		elfPairs[i] = ElfPair{firstElf: firstElf, secondElf: secondElf}
	}

	return elfPairs
}

func getElfFromRangeString(elfRangeString string) Elf {
	elfRange := strings.Split(elfRangeString, "-")

	sectionStart, _ := strconv.Atoi(elfRange[0])
	sectionEnd, _ := strconv.Atoi(elfRange[1])

	return Elf{sectionStart: sectionStart, sectionEnd: sectionEnd}
}

func (Day4 Day4) FirstTaskResult() {
	println("DAY4 first task")

	elfPairs := createElfPairsFromInput(Day4)

	totalContains := 0
	for _, pair := range elfPairs {
		if pair.rangeFullyContains() {
			totalContains++
		}
	}

	fmt.Println(totalContains)
}

func (Day4 Day4) SecondTaskResult() {
	println("DAY4 second task")

	elfPairs := createElfPairsFromInput(Day4)

	totalContains := 0
	for _, pair := range elfPairs {
		if pair.rangeContains() {
			totalContains++
		}
	}

	fmt.Println(totalContains)
}

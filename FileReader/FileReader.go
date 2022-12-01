package FileReader

import (
	"AOC_2022/ErrorHandler"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getFileScanner(filePath string) *bufio.Scanner {
	fmt.Println("Opening:" + filePath)
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Failed to open file")
		log.Fatal(err)
	}

	return bufio.NewScanner(file)
}

func GetIntMap(filePath string) map[int]int {
	m := make(map[int]int)
	i := 0

	scanner := getFileScanner(filePath)

	for scanner.Scan() {
		m[getLineAsInt(scanner.Text())] = getLineAsInt(scanner.Text())

		i++
	}

	return m
}

func GetIntArray(filePath string) []int {
	scanner := getFileScanner(filePath)

	var values []int

	for scanner.Scan() {
		values = append(values, getLineAsInt(scanner.Text()))
	}

	return values
}

func getLineAsInt(line string) int {
	if line == "" {
		return 0
	}

	intVar, err := strconv.Atoi(line)
	ErrorHandler.Handle(err)

	return intVar
}

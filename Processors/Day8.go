package Processors

import (
	"AOC_2022/FileReader"
	"fmt"
	"strconv"
	"strings"
)

type Day8 struct {
	Processor
	InputFilePath string
}

func getTreeMap(Day Day8) (map[string]int, int, int) {
	inputLines := FileReader.GetStringArray(Day.InputFilePath)
	trees := make(map[string]int)

	for y, line := range inputLines {
		treeRow := strings.Split(line, "")
		for x, treeSizeString := range treeRow {
			treeSize, _ := strconv.Atoi(treeSizeString)
			key := getCoordKey(x, y)
			trees[key] = treeSize
		}
	}

	rows := len(inputLines)
	columns := len(strings.Split(inputLines[0], ""))

	return trees, columns, rows
}

func (Day Day8) FirstTaskResult() {
	treeMap, rows, columns := getTreeMap(Day)

	isVisibleLeft(1, 1, treeMap, columns, rows)

	visibleTrees := 0
	for y := 0; y < rows; y++ {
		for x := 0; x < columns; x++ {
			if isTreeVisible(x, y, treeMap, columns, rows) {
				visibleTrees++
			}
		}
	}

	fmt.Println(visibleTrees)
}

func isTreeVisible(treeXCord int, treeYCord int, treeMap map[string]int, columns int, rows int) bool {
	if isVisibleUp(treeXCord, treeYCord, treeMap, columns, rows) {
		println("IS VISIBLE UP " + getCoordKey(treeXCord, treeYCord))
		return true
	}

	if isVisibleDown(treeXCord, treeYCord, treeMap, columns, rows) {
		println("IS VISIBLE DOWN " + getCoordKey(treeXCord, treeYCord))
		return true
	}

	if isVisibleRight(treeXCord, treeYCord, treeMap, columns, rows) {
		println("IS VISIBLE RIGHT " + getCoordKey(treeXCord, treeYCord))
		return true
	}

	if isVisibleLeft(treeXCord, treeYCord, treeMap, columns, rows) {
		println("IS VISIBLE LEFT " + getCoordKey(treeXCord, treeYCord))
		return true
	}

	println("IS NOT VISIBLE " + getCoordKey(treeXCord, treeYCord))
	return false
}

func isVisibleDown(x int, y int, treeMap map[string]int, columns int, rows int) bool {
	currentTreeSize := treeMap[getCoordKey(x, y)]
	//currentTreeCordKey := getCoordKey(x, y)

	for y := y + 1; y < columns; y++ {
		treeToCheckSize := treeMap[getCoordKey(x, y)]

		if treeToCheckSize >= currentTreeSize {
			//fmt.Println(currentTreeCordKey + " :DOWN _LARGER/SAME FOUND at:" + getCoordKey(x, y))
			return false
		}
	}

	//fmt.Println(currentTreeCordKey + ":DOWN NOT FOUND")
	return true
}

func isVisibleUp(x int, y int, treeMap map[string]int, columns int, rows int) bool {
	currentTreeSize := treeMap[getCoordKey(x, y)]
	//currentTreeCordKey := getCoordKey(x, y)

	for y := y + -1; y >= 0; y-- {
		treeToCheckSize := treeMap[getCoordKey(x, y)]

		if treeToCheckSize >= currentTreeSize {
			//fmt.Println(currentTreeCordKey + ":UP :LARGER/SAME FOUND at:" + getCoordKey(x, y))
			return false
		}
	}

	//fmt.Println(currentTreeCordKey + "UP Larger NOT FOUND")
	return true
}

func isVisibleRight(x int, y int, treeMap map[string]int, columns int, rows int) bool {
	currentTreeSize := treeMap[getCoordKey(x, y)]
	//currentTreeCordKey := getCoordKey(x, y)

	for x := x + 1; x < rows; x++ {
		treeToCheckSize := treeMap[getCoordKey(x, y)]

		if treeToCheckSize >= currentTreeSize {
			//fmt.Println(currentTreeCordKey + ": -->>RIGHT LARGER/SAME FOUND at:" + getCoordKey(x, y))
			return false
		}
	}

	//fmt.Println(currentTreeCordKey + " :Larger NOT FOUND")
	return true
}

func isVisibleLeft(x int, y int, treeMap map[string]int, columns int, rows int) bool {
	currentTreeSize := treeMap[getCoordKey(x, y)]
	//currentTreeCordKey := getCoordKey(x, y)

	//fmt.Println("___DOWN___:" + currentTreeCordKey)

	for x := x - 1; x >= 0; x-- {
		treeToCheckSize := treeMap[getCoordKey(x, y)]

		if treeToCheckSize >= currentTreeSize {
			//fmt.Println("LARGER/SAME FOUND at:" + getCoordKey(x, y))
			return false
		}
	}

	//fmt.Println("Larger NOT FOUND")
	return true
}

func getXY(string string) (int, int) {
	split := strings.Split(string, "_")

	x, _ := strconv.Atoi(split[0])
	y, _ := strconv.Atoi(split[1])

	return x, y
}

func getCoordKey(x int, y int) string {
	xS := strconv.Itoa(x)
	yS := strconv.Itoa(y)

	return xS + "_" + yS
}

func (Day Day8) SecondTaskResult() {

}

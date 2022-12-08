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

	visibleTrees := 0
	bestScenicScore := 0
	for y := 0; y < rows; y++ {
		for x := 0; x < columns; x++ {
			isVisible, scenicScore := isTreeVisible(x, y, treeMap, columns, rows)

			if scenicScore > bestScenicScore {
				bestScenicScore = scenicScore
			}

			if isVisible {
				visibleTrees++
			}
		}
	}

	fmt.Println(visibleTrees)
	fmt.Println(bestScenicScore)
}

func isTreeVisible(treeXCord int, treeYCord int, treeMap map[string]int, columns int, rows int) (bool, int) {

	isVisibleUp, visibilityUpLenght := isVisibleUp(treeXCord, treeYCord, treeMap, columns, rows)
	isVisibleDown, visiblityDownLenght := isVisibleDown(treeXCord, treeYCord, treeMap, columns, rows)
	isVisibleRight, visiblityRightLenght := isVisibleRight(treeXCord, treeYCord, treeMap, columns, rows)
	isVisibleLeft, visibilityLeftLeght := isVisibleLeft(treeXCord, treeYCord, treeMap, columns, rows)

	isVisible := isVisibleUp || isVisibleDown || isVisibleLeft || isVisibleRight
	scenicScore := visibilityUpLenght * visiblityDownLenght * visibilityLeftLeght * visiblityRightLenght

	return isVisible, scenicScore
}

func isVisibleDown(x int, y int, treeMap map[string]int, columns int, rows int) (bool, int) {
	currentTreeSize := treeMap[getCoordKey(x, y)]
	//currentTreeCordKey := getCoordKey(x, y)
	visibilityLenght := 0
	for y := y + 1; y < columns; y++ {
		treeToCheckSize := treeMap[getCoordKey(x, y)]
		visibilityLenght++

		if treeToCheckSize >= currentTreeSize {
			//fmt.Println(currentTreeCordKey + " :DOWN _LARGER/SAME FOUND at:" + getCoordKey(x, y))
			return false, visibilityLenght
		}
	}

	//fmt.Println(currentTreeCordKey + ":DOWN NOT FOUND")
	return true, visibilityLenght
}

func isVisibleUp(x int, y int, treeMap map[string]int, columns int, rows int) (bool, int) {
	currentTreeSize := treeMap[getCoordKey(x, y)]
	//currentTreeCordKey := getCoordKey(x, y)

	visibilityLenght := 0
	for y := y + -1; y >= 0; y-- {
		treeToCheckSize := treeMap[getCoordKey(x, y)]
		visibilityLenght++
		if treeToCheckSize >= currentTreeSize {
			//fmt.Println(currentTreeCordKey + ":UP :LARGER/SAME FOUND at:" + getCoordKey(x, y))
			return false, visibilityLenght
		}
	}

	//fmt.Println(currentTreeCordKey + "UP Larger NOT FOUND")
	return true, visibilityLenght
}

func isVisibleRight(x int, y int, treeMap map[string]int, columns int, rows int) (bool, int) {
	currentTreeSize := treeMap[getCoordKey(x, y)]
	//currentTreeCordKey := getCoordKey(x, y)

	visibilityLenght := 0

	for x := x + 1; x < rows; x++ {
		treeToCheckSize := treeMap[getCoordKey(x, y)]
		visibilityLenght++

		if treeToCheckSize >= currentTreeSize {
			//fmt.Println(currentTreeCordKey + ": -->>RIGHT LARGER/SAME FOUND at:" + getCoordKey(x, y))
			return false, visibilityLenght
		}
	}

	//fmt.Println(currentTreeCordKey + " :Larger NOT FOUND")
	return true, visibilityLenght
}

func isVisibleLeft(x int, y int, treeMap map[string]int, columns int, rows int) (bool, int) {
	currentTreeSize := treeMap[getCoordKey(x, y)]
	//currentTreeCordKey := getCoordKey(x, y)

	//fmt.Println("___DOWN___:" + currentTreeCordKey)

	visibilityLenght := 0

	for x := x - 1; x >= 0; x-- {
		treeToCheckSize := treeMap[getCoordKey(x, y)]
		visibilityLenght++
		if treeToCheckSize >= currentTreeSize {
			//fmt.Println("LARGER/SAME FOUND at:" + getCoordKey(x, y))
			return false, visibilityLenght
		}
	}

	//fmt.Println("Larger NOT FOUND")
	return true, visibilityLenght
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

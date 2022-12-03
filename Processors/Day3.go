package Processors

//Dont judge - hangover stuff O(n^infinity) - but go is fast so didn't have to wait all day
import (
	"AOC_2022/FileReader"
	"fmt"
	"strings"
	"unicode"
)

type Day3 struct {
	InputFilePath string
	processor     Processor
}

func (Day3 Day3) FirstTaskResult() {
	inputLines := FileReader.GetStringArray(Day3.InputFilePath)

	//fmt.Println(getPointsForCharacter("B"))
	//fmt.Println(getPointsForCharacter("b"))
	//fmt.Println([]rune("A")[0])
	//fmt.Println([]rune("Z")[0])

	totalPoints := 0
	for _, inputLine := range inputLines {

		for _, points := range itemsAndPriorityInBothCompartments(inputLine) {
			totalPoints += points
		}
	}

	fmt.Println(totalPoints)

}

func (Day3 Day3) SecondTaskResult() {
	itemsByElf := FileReader.GetStringArray(Day3.InputFilePath)

	itemsByGroup := getItemsByElfGroup(itemsByElf)

	totalPoints := 0
	for _, itemsByElfGroup := range itemsByGroup {
		fmt.Println(itemsByElfGroup)
		itemInAllsacks := getDuplicateItemInElfSacks(itemsByElfGroup)

		totalPoints += getPointsForCharacter(itemInAllsacks)
	}

	print(totalPoints)
}

func getDuplicateItemInElfSacks(itemsByElfGroup [3]string) string {
	firstSack := itemsByElfGroup[0]
	secondSack := itemsByElfGroup[1]
	thirdSack := itemsByElfGroup[2]

	for _, firstSackItem := range strings.Split(firstSack, "") {
		for _, secondSackItem := range strings.Split(secondSack, "") {
			for _, thirdSackItem := range strings.Split(thirdSack, "") {
				if firstSackItem == secondSackItem {
					if secondSackItem == thirdSackItem {
						println("In all sacks:" + firstSackItem)
						return firstSackItem
					}
				}
			}
		}
	}

	return ""
}

func getItemsByElfGroup(itemsByElf []string) map[int][3]string {
	elfIndex := 0
	groupIndex := 0

	itemsByGroup := make(map[int][3]string, len(itemsByElf)/3)
	for i, _ := range itemsByElf {
		if (i+1)%3 == 0 {
			var itemsInGroup [3]string
			itemsInGroup[0] = itemsByElf[i-2]
			itemsInGroup[1] = itemsByElf[i-1]
			itemsInGroup[2] = itemsByElf[i]
			itemsByGroup[groupIndex] = itemsInGroup
			groupIndex++
		}
		elfIndex++
	}

	return itemsByGroup
}

func itemsAndPriorityInBothCompartments(inputLine string) map[string]int {
	characterArray := strings.Split(inputLine, "")
	sizeOfString := len(characterArray)

	firstCompartment := characterArray[0 : sizeOfString/2]
	secondCompartment := characterArray[sizeOfString/2 : sizeOfString]

	itemsInBothCompartments := make(map[string]int)
	for _, x := range firstCompartment {
		for _, y := range secondCompartment {
			if x == y {
				itemsInBothCompartments[x] = getPointsForCharacter(x)
			}
		}
	}

	return itemsInBothCompartments
}

func getPointsForCharacter(character string) int {
	//a-z 97 - 122
	//A-Z 65 - 90
	number := int([]rune(character)[0])

	if IsUpper(character) {
		return number - 38

	}
	return number - 96
}

// Stackoverflow
func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

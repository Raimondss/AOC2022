package Processors

//MAde while dunk
import (
	fileReader "AOC_2022/FileReader"
	"fmt"
	"strings"
)

type Day2 struct {
	InputFilePath string
	processor     Processor
}

const ResultWin = "WIN"
const ResultLoose = "LOOSE"
const ResultDraw = "DRAW"
const HandScissors = "scissors"
const HandRock = "rock"
const HandPaper = "paper"

// Hand -> hand that wins it
var winningHandMap = map[string]string{
	HandRock:     HandPaper,
	HandScissors: HandRock,
	HandPaper:    HandScissors,
}

var LosingHandMap = map[string]string{
	HandRock:     HandScissors,
	HandScissors: HandPaper,
	HandPaper:    HandRock,
}

var inputResultMap = map[string]string{
	"X": ResultLoose,
	"Y": ResultDraw,
	"Z": ResultWin,
}
var resultPointMap = map[string]int{
	ResultLoose: 0,
	ResultDraw:  3,
	ResultWin:   6,
}

var handPointMap = map[string]int{
	HandRock:     1,
	HandPaper:    2,
	HandScissors: 3,
}

var inputHandMap = map[string]string{
	"X": HandRock,
	"Y": HandPaper,
	"Z": HandScissors,
	"A": HandRock,
	"B": HandPaper,
	"C": HandScissors,
}

func (Day2 Day2) FirstTaskResult() {
	fmt.Println("Day 2-1 task")
	stringArray := fileReader.GetStringArray(Day2.InputFilePath)

	totalPoints := 0
	for i := 0; i < len(stringArray); i++ {
		inputLine := stringArray[i]

		split := strings.Split(inputLine, " ")
		theirHand := inputHandMap[split[0]]
		myHand := inputHandMap[split[1]]

		result := getGameResult(theirHand, myHand)
		pointsInGame := handPointMap[myHand] + resultPointMap[result]
		totalPoints += pointsInGame
	}

	println(totalPoints)
}

func getGameResult(theirHand string, myHand string) string {
	if theirHand == myHand {
		return ResultDraw
	}

	if winningHandMap[theirHand] == myHand {
		return ResultWin
	}

	return ResultLoose
}

func handForPreferredResult(theirHand string, preferredResult string) string {
	if preferredResult == ResultDraw {
		return theirHand
	}

	if preferredResult == ResultWin {
		return winningHandMap[theirHand]
	}

	return LosingHandMap[theirHand]
}

func (Day2 Day2) SecondTaskResult() {
	stringArray := fileReader.GetStringArray(Day2.InputFilePath)

	totalPoints := 0
	for i := 0; i < len(stringArray); i++ {
		inputLine := stringArray[i]
		split := strings.Split(inputLine, " ")
		theirHand := inputHandMap[split[0]]
		preferredResult := inputResultMap[split[1]]

		myHand := handForPreferredResult(theirHand, preferredResult)

		result := getGameResult(theirHand, myHand)
		pointsInGame := handPointMap[myHand] + resultPointMap[result]
		totalPoints += pointsInGame
	}

	println(totalPoints)
}

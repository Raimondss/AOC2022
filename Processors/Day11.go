package Processors

import (
	"fmt"
	"math"
	"strconv"
)

type Day11 struct {
	Processor
	InputFilePath string
}

type Game struct {
}

type Operation func(int) int

type Monkey struct {
	holdingItems       map[int]int
	inspectedItemCount int
	operation          Operation
	testInt            int
	trueReceiver       int
	falseReceiver      int
	index              int
}

func (Monkey Monkey) inspectItems(monkeyArray map[int]Monkey) map[int]Monkey {
	fmt.Println("Monkey " + strconv.Itoa(Monkey.index) + " begins items inspection")

	for itemIndex, itemWorryLevel := range Monkey.holdingItems {
		fmt.Println("Inspects:" + strconv.Itoa(itemWorryLevel))
		Monkey.inspectedItemCount++
		itemWorryLevel = Monkey.operation(itemWorryLevel)
		fmt.Println("Worry level changed to:" + strconv.Itoa(itemWorryLevel))

		itemWorryLevel = int(math.Round(float64(itemWorryLevel / 3)))
		fmt.Println("Worry level changed to:" + strconv.Itoa(itemWorryLevel))

		isDivisible := itemWorryLevel%Monkey.testInt == 0

		monkeyToThrowTo := 0

		if isDivisible {
			fmt.Println("Is divisible by " + strconv.Itoa(Monkey.testInt) + " thrown to:" + strconv.Itoa(Monkey.trueReceiver))
			monkeyToThrowTo = Monkey.trueReceiver
		} else {
			fmt.Println("Is not divisible by " + strconv.Itoa(Monkey.testInt) + " thrown to:" + strconv.Itoa(Monkey.falseReceiver))
			monkeyToThrowTo = Monkey.falseReceiver
		}

		monkeyToReceive := monkeyArray[monkeyToThrowTo]
		monkeyToReceive.addItem(itemWorryLevel)
		Monkey.removeItem(itemIndex)

		monkeyArray[Monkey.index] = Monkey
	}

	return monkeyArray
}

func (Monkey *Monkey) addItem(item int) {
	Monkey.holdingItems[len(Monkey.holdingItems)+1] = item
}

func (Monkey *Monkey) removeItem(index int) {
	delete(Monkey.holdingItems, index)
}

func (Day Day11) FirstTaskResult() {
	fmt.Println("First")

	monkeys := createMonkeys()

	for round := 1; round <= 20; round++ {
		fmt.Println("ROUND:" + strconv.Itoa(round))
		fmt.Println("$$$$$$$$$$$$$$")
		for i := 0; i < len(monkeys); i++ {
			fmt.Println("_____________")
			currentMonkey := monkeys[i]
			monkeys = currentMonkey.inspectItems(monkeys)
		}
	}
	println("________TASK1____________")
	for _, monkey := range monkeys {
		fmt.Println("Monkey" + strconv.Itoa(monkey.index) + " inspected items " + strconv.Itoa(monkey.inspectedItemCount))
	}

}

func (Day Day11) SecondTaskResult() {
	fmt.Println("Second")
	monkeys := createMonkeys()

	dividor := 1
	for _, monkey := range monkeys {
		dividor *= monkey.testInt
	}

	for round := 1; round <= 10000; round++ {
		//fmt.Println("ROUND:" + strconv.Itoa(round))
		//fmt.Println("$$$$$$$$$$$$$$")
		for i := 0; i < len(monkeys); i++ {
			//fmt.Println("_____________")
			currentMonkey := monkeys[i]
			monkeys = currentMonkey.inspectItemsTwo(monkeys, dividor)
		}

		if round%1000 == 0 {
			println(round)
			for _, monkey := range monkeys {
				fmt.Println("Monkey" + strconv.Itoa(monkey.index) + " inspected items " + strconv.Itoa(monkey.inspectedItemCount))
			}
		}
	}

	println("RESULT")
	for _, monkey := range monkeys {
		fmt.Println("Monkey" + strconv.Itoa(monkey.index) + " inspected items " + strconv.Itoa(monkey.inspectedItemCount))
	}
}

func (Monkey Monkey) inspectItemsTwo(monkeyArray map[int]Monkey, commonDividor int) map[int]Monkey {
	//fmt.Println("Monkey " + strconv.Itoa(Monkey.index) + " begins items inspection")

	for itemIndex, itemWorryLevel := range Monkey.holdingItems {
		//fmt.Println("Inspects:" + strconv.Itoa(itemWorryLevel))
		if itemWorryLevel > commonDividor {
			itemWorryLevel = itemWorryLevel % commonDividor
		}

		Monkey.inspectedItemCount++
		itemWorryLevel = Monkey.operation(itemWorryLevel)
		//fmt.Println("Worry level changed to:" + strconv.Itoa(itemWorryLevel))

		isDivisible := itemWorryLevel%Monkey.testInt == 0

		monkeyToThrowTo := 0

		if isDivisible {
			//fmt.Println("Is divisible by " + strconv.Itoa(Monkey.testInt) + " thrown to:" + strconv.Itoa(Monkey.trueReceiver))
			monkeyToThrowTo = Monkey.trueReceiver
		} else {
			//fmt.Println("Is not divisible by " + strconv.Itoa(Monkey.testInt) + " thrown to:" + strconv.Itoa(Monkey.falseReceiver))
			monkeyToThrowTo = Monkey.falseReceiver
		}

		monkeyToReceive := monkeyArray[monkeyToThrowTo]
		monkeyToReceive.addItem(itemWorryLevel)
		Monkey.removeItem(itemIndex)

		monkeyArray[Monkey.index] = Monkey
	}

	return monkeyArray
}

func createMonkeys() map[int]Monkey {

	var monkeys = make(map[int]Monkey, 8)

	var monkey0 = Monkey{
		falseReceiver: 1,
		trueReceiver:  7,
		operation: func(worryLevel int) int {
			return worryLevel * 7
		},
		testInt:            2,
		inspectedItemCount: 0,
		holdingItems:       map[int]int{1: 62, 2: 92, 3: 50, 4: 63, 5: 62, 6: 93, 7: 73, 8: 50},
		index:              0,
	}

	monkeys[0] = monkey0

	var monkey1 = Monkey{
		holdingItems:  map[int]int{1: 51, 2: 97, 3: 74, 4: 84, 5: 99},
		trueReceiver:  2,
		falseReceiver: 4,
		operation: func(worryLevel int) int {
			return worryLevel + 3
		},
		testInt:            7,
		inspectedItemCount: 0,
		index:              1,
	}

	monkeys[1] = monkey1

	var monkey2 = Monkey{
		holdingItems:  map[int]int{1: 98, 2: 86, 3: 62, 4: 76, 5: 51, 6: 81, 7: 95},
		trueReceiver:  5,
		falseReceiver: 4,
		operation: func(worryLevel int) int {
			return worryLevel + 4
		},
		testInt:            13,
		inspectedItemCount: 0,
		index:              2,
	}

	monkeys[2] = monkey2

	var monkey3 = Monkey{
		holdingItems:  map[int]int{1: 53, 2: 95, 3: 50, 4: 85, 5: 83, 6: 72},
		trueReceiver:  6,
		falseReceiver: 0,
		operation: func(worryLevel int) int {
			return worryLevel + 5
		},
		testInt:            19,
		inspectedItemCount: 0,
		index:              3,
	}

	monkeys[3] = monkey3

	var monkey4 = Monkey{
		holdingItems:  map[int]int{1: 59, 2: 60, 3: 63, 4: 71},
		trueReceiver:  5,
		falseReceiver: 3,
		operation: func(worryLevel int) int {
			return worryLevel * 5
		},
		testInt:            11,
		inspectedItemCount: 0,
		index:              4,
	}

	monkeys[4] = monkey4

	var monkey5 = Monkey{
		holdingItems:  map[int]int{1: 92, 2: 65},
		trueReceiver:  6,
		falseReceiver: 3,
		operation: func(worryLevel int) int {
			return worryLevel * worryLevel
		},
		testInt:            5,
		inspectedItemCount: 0,
		index:              5,
	}

	monkeys[5] = monkey5

	var monkey6 = Monkey{
		holdingItems:  map[int]int{1: 78},
		trueReceiver:  0,
		falseReceiver: 7,
		operation: func(worryLevel int) int {
			return worryLevel + 8
		},
		testInt:            3,
		inspectedItemCount: 0,
		index:              6,
	}

	monkeys[6] = monkey6

	var monkey7 = Monkey{
		holdingItems:  map[int]int{1: 84, 2: 93, 3: 54},
		trueReceiver:  2,
		falseReceiver: 1,
		operation: func(worryLevel int) int {
			return worryLevel + 1
		},
		testInt:            17,
		inspectedItemCount: 0,
		index:              7,
	}

	monkeys[7] = monkey7

	return monkeys
}

func createTestMonkeys() map[int]Monkey {
	var monkeys = make(map[int]Monkey, 4)

	var monkey0 = Monkey{
		trueReceiver:  2,
		falseReceiver: 3,
		operation: func(worryLevel int) int {
			return worryLevel * 19
		},
		testInt:            23,
		inspectedItemCount: 0,
		holdingItems:       map[int]int{1: 79, 2: 98},
		index:              0,
	}

	monkeys[0] = monkey0

	var monkey1 = Monkey{
		holdingItems:  map[int]int{1: 54, 2: 65, 3: 75, 4: 74},
		trueReceiver:  2,
		falseReceiver: 0,
		operation: func(worryLevel int) int {
			return worryLevel + 6
		},
		testInt:            19,
		inspectedItemCount: 0,
		index:              1,
	}

	monkeys[1] = monkey1

	var monkey2 = Monkey{
		holdingItems:  map[int]int{1: 79, 2: 60, 3: 97},
		trueReceiver:  1,
		falseReceiver: 3,
		operation: func(worryLevel int) int {
			return worryLevel * worryLevel
		},
		testInt:            13,
		inspectedItemCount: 0,
		index:              2,
	}

	monkeys[2] = monkey2

	var monkey3 = Monkey{
		holdingItems:  map[int]int{1: 74},
		trueReceiver:  0,
		falseReceiver: 1,
		operation: func(worryLevel int) int {
			return worryLevel + 3
		},
		testInt:            17,
		inspectedItemCount: 0,
		index:              3,
	}

	monkeys[3] = monkey3

	return monkeys
}

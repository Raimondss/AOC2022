package Processors

import (
	"AOC_2022/FileReader"
	"fmt"
	"strconv"
	"strings"
)

type Day9 struct {
	Processor
	InputFilePath string
}

type Movement struct {
	direction string
	times     int
}

type Position struct {
	x int
	y int
}

func (Position Position) getPositionKey() string {
	return strconv.Itoa(Position.x) + "_" + strconv.Itoa(Position.y)
}

type Tile struct {
	position               Position
	uniquePositionsVisited map[string]bool
}

func (Rope *Rope) checkTail() {
	diffX := Rope.head.position.x - Rope.tail.position.x
	diffY := Rope.head.position.y - Rope.tail.position.y

	if diffX <= 1 && diffX >= -1 && diffY <= 1 && diffY >= -1 {
		//In connection
		return
	}

	direction := ""

	if diffX == 2 {

		if diffY == 0 {
			direction = DirectionRight
		}

		if diffY == 1 {
			direction = DirectionUpRight
		}

		if diffY == -1 {
			direction = DirectionDownRight
		}
	}

	if diffX == -2 {
		if diffY == 0 {
			direction = DirectionLeft
		}

		if diffY == 1 {
			direction = DirectionUpLeft
		}

		if diffY == -1 {
			direction = DirectionDownLeft
		}
	}

	if diffY == 2 {
		if diffX == 0 {
			direction = DirectionUp
		}

		if diffX == 1 {
			direction = DirectionUpRight
		}

		if diffX == -1 {
			direction = DirectionUpLeft
		}
	}

	if diffY == -2 {
		if diffX == 0 {
			direction = DirectionDown
		}

		if diffX == 1 {
			direction = DirectionDownRight
		}

		if diffX == -1 {
			direction = DirectionDownLeft
		}
	}

	if direction == "" {
		panic("PANIC - NO FOLLOW DIRECTION")
	}

	newTailPosition := getNewPositionBasedOnDirection(direction, Rope.tail.position)

	Rope.tail.setPosition(newTailPosition)
}

type Rope struct {
	head Tile
	tail Tile
}

func (Rope *Rope) doMovement(Movement Movement) {
	for i := 0; i < Movement.times; i++ {
		Rope.moveHeadByDirection(Movement.direction)
	}
}

func (Rope *Rope) moveHeadByDirection(direction string) {
	newPosition := getNewPositionBasedOnDirection(direction, Rope.head.position)
	Rope.head.setPosition(newPosition)
	Rope.checkTail()
}

const DirectionUp = "U"
const DirectionDown = "D"
const DirectionRight = "R"
const DirectionLeft = "L"
const DirectionUpLeft = "UL"
const DirectionUpRight = "UR"
const DirectionDownLeft = "DL"
const DirectionDownRight = "DR"

func getNewPositionBasedOnDirection(direction string, position Position) Position {
	if direction == DirectionRight {
		position.x += 1
		return position
	}

	if direction == DirectionLeft {
		position.x += -1
		return position
	}

	if direction == DirectionDown {
		position.y += -1
		return position
	}

	if direction == DirectionUp {
		position.y += +1

		return position
	}

	if direction == DirectionUpRight {
		position.x += 1
		position.y += 1

		return position
	}

	if direction == DirectionUpLeft {
		position.x += -1
		position.y += 1

		return position
	}

	if direction == DirectionDownRight {
		position.x += 1
		position.y += -1

		return position
	}

	if direction == DirectionDownLeft {
		position.x += -1
		position.y += -1

		return position
	}

	panic("PANIC - MOVEMENT NOT VALID")
}

func (Tile *Tile) setPosition(newPosition Position) {
	Tile.uniquePositionsVisited[newPosition.getPositionKey()] = true

	Tile.position = newPosition
}

func (Movement Movement) getMovementAxis() string {
	if Movement.direction == "R" || Movement.direction == "L" {
		return "X"
	}

	return "Y"
}

func (Movement Movement) getMovementSign() int {
	if Movement.direction == "D" || Movement.direction == "L" {
		return -1
	}

	return 1
}

func (Day9 Day9) FirstTaskResult() {
	fmt.Print("First:")
	headMovements := getHeadMovements(Day9)

	rope := Rope{head: Tile{
		position:               Position{x: 0, y: 0},
		uniquePositionsVisited: make(map[string]bool),
	},
		tail: Tile{
			position:               Position{x: 0, y: 0},
			uniquePositionsVisited: make(map[string]bool),
		}}

	rope.tail.uniquePositionsVisited[rope.tail.position.getPositionKey()] = true
	rope.head.uniquePositionsVisited[rope.tail.position.getPositionKey()] = true

	for _, headMovement := range headMovements {
		rope.doMovement(headMovement)
	}

	fmt.Println(len(rope.tail.uniquePositionsVisited))

}

func getHeadMovements(Day9 Day9) []Movement {
	movementCommand := FileReader.GetStringArray(Day9.InputFilePath)

	movements := make([]Movement, len(movementCommand))
	for i, s := range movementCommand {
		split := strings.Split(s, " ")
		times, _ := strconv.Atoi(split[1])
		direction := split[0]

		command := Movement{direction: direction, times: times}

		movements[i] = command
	}

	return movements
}

func (Day9 Day9) SecondTaskResult() {

}

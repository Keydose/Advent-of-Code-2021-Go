package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Moveable interface {
	move(deltaX int, deltaY int)
	Position() Coordinates
}

type Coordinates struct {
	x int
	y int
}

type Submarine struct {
	position Coordinates
}

func (s *Submarine) move(deltaCoords Coordinates) {
	s.position.x += deltaCoords.x
	s.position.y += deltaCoords.y
}

func (s Submarine) Position() Coordinates {
	return s.position
}

func parseMovementInstruction(instruction string) Coordinates {
	instructionParts := strings.Split(instruction, " ")
	direction := instructionParts[0]
	magnitude, err := strconv.Atoi(instructionParts[1])
	if err != nil {
		log.Fatal(err)
	}

	if direction == "up" {
		return Coordinates{x: 0, y: -magnitude}
	} else if direction == "down" {
		return Coordinates{x: 0, y: magnitude}
	} else if direction == "forward" {
		return Coordinates{x: magnitude, y: 0}
	} else {
		// TODO: Figure out how to throw an exception instead?
		return Coordinates{x: 0, y: 0}
	}
}

func main() {
	file, err := os.Open("movement_instructions.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	submarine := Submarine{position: Coordinates{x: 0, y: 0}}
	for scanner.Scan() {
		instruction := scanner.Text()
		submarine.move(parseMovementInstruction(instruction))
	}

	submarinePosition := submarine.Position()

	fmt.Println(submarinePosition.x * submarinePosition.y)
}

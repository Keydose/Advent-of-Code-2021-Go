package main

import (
	"bufio"
	"errors"
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
	aim      int
}

func (s *Submarine) move(deltaCoords Coordinates) {
	if deltaCoords.x != 0 {
		s.position.x += deltaCoords.x
		if s.aim > 0 {
			s.position.y += deltaCoords.x * s.aim
		}
	} else if deltaCoords.y != 0 {
		s.aim += deltaCoords.y
	}
}

func (s Submarine) Position() Coordinates {
	return s.position
}

func parseMovementInstruction(instruction string) (*Coordinates, error) {
	instructionParts := strings.Split(instruction, " ")
	direction := instructionParts[0]
	magnitude, err := strconv.Atoi(instructionParts[1])
	if err != nil {
		return nil, errors.New("magnitude not an integer in instruction: " + instruction)
	}

	if direction == "up" {
		return &Coordinates{x: 0, y: -magnitude}, nil
	} else if direction == "down" {
		return &Coordinates{x: 0, y: magnitude}, nil
	} else if direction == "forward" {
		return &Coordinates{x: magnitude, y: 0}, nil
	} else {
		return nil, errors.New("unrecognised direction in instruction: " + instruction)
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
		movementInstruction, err := parseMovementInstruction(scanner.Text())
		if err != nil {
			fmt.Println(err)
			continue
		}
		submarine.move(*movementInstruction)
	}

	submarinePosition := submarine.Position()

	fmt.Println(submarinePosition.x * submarinePosition.y)
}

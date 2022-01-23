package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}

func moveSubmarine(instruction string, submarinePosition *Position) {
	instructionParts := strings.Split(instruction, " ")
	direction := instructionParts[0]
	magnitude, err := strconv.Atoi(instructionParts[1])
	if err != nil {
		log.Fatal(err)
	}

	if direction == "up" {
		submarinePosition.y -= magnitude
	} else if direction == "down" {
		submarinePosition.y += magnitude
	} else if direction == "forward" {
		submarinePosition.x += magnitude
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	submarinePosition := Position{x: 0, y: 0}
	for scanner.Scan() {
		instruction := scanner.Text()
		moveSubmarine(instruction, &submarinePosition)
	}

	fmt.Println(submarinePosition.x * submarinePosition.y)
}

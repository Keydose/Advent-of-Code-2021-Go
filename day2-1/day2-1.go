package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	deltaX := 0
	deltaY := 0
	for scanner.Scan() {
		instruction := scanner.Text()

		instructionParts := strings.Split(instruction, " ")
		direction := instructionParts[0]
		magnitude, err := strconv.Atoi(instructionParts[1])
		if err != nil {
			log.Fatal(err)
		}

		if direction == "up" {
			deltaY -= magnitude
		} else if direction == "down" {
			deltaY += magnitude
		} else if direction == "forward" {
			deltaX += magnitude
		}
	}

	fmt.Println(deltaX * deltaY)
}

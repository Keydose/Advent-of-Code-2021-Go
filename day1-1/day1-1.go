package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var increasedCount int = 0
	var previousValue int = 0
	var firstValue bool = true
	for scanner.Scan() {
		currentValue, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		if !firstValue {
			if currentValue > previousValue {
				increasedCount++
			}
		} else {
			firstValue = false
		}

		previousValue = currentValue
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(increasedCount)
}

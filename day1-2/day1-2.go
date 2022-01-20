package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	return sum
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	measurementWindow := []int{0, 0, 0}
	var previousWindowSum int = 0
	var latestWindowSum int = 0
	var increasedCount int = 0

	var measurementsCount int = 0

	for scanner.Scan() {
		latestMeasurement, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		measurementsCount++

		previousWindowSum = sum(measurementWindow)
		measurementWindow = append(measurementWindow[1:], latestMeasurement)
		latestWindowSum = sum(measurementWindow)

		if measurementsCount > 3 {
			if latestWindowSum > previousWindowSum {
				increasedCount++
			}
		}
	}

	fmt.Println(increasedCount)
}

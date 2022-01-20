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
		// Chop off the first element of slice, and add the latest measurement to the end
		// (so slice size of 3 never changes)
		measurementWindow = append(measurementWindow[1:], latestMeasurement)
		latestWindowSum = sum(measurementWindow)

		// Sliding window is 3 measurements wide
		// We can only start comparing windows once we have had a full window
		if measurementsCount > 3 {
			if latestWindowSum > previousWindowSum {
				increasedCount++
			}
		}
	}

	fmt.Println(increasedCount)
}

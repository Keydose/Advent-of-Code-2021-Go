package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("diagnostic_report.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	diagnosticCount := 0
	diagnosticLength := 0
	var diagnosticBitSums []int
	for scanner.Scan() {
		diagnosticBits := []rune(scanner.Text())
		if diagnosticCount == 0 {
			// First diagnostic - create a slice with matching length to store
			// sum of occurrences of '1' in each place
			diagnosticLength = len(diagnosticBits)
			diagnosticBitSums = make([]int, diagnosticLength)
		}

		for i := 0; i < diagnosticLength; i++ {
			if diagnosticBits[i] == '1' {
				diagnosticBitSums[i]++
			}
		}

		diagnosticCount++
	}

	gammaRateBits := make([]rune, diagnosticLength)
	epsilonRateBits := make([]rune, diagnosticLength)

	for i := 0; i < diagnosticLength; i++ {
		if diagnosticBitSums[i] > diagnosticCount/2 {
			// 1 is the most common in i'th place
			gammaRateBits[i] = '1'
			epsilonRateBits[i] = '0'
		} else {
			// 0 is the most common in i'th place
			gammaRateBits[i] = '0'
			epsilonRateBits[i] = '1'
		}
	}

	// Convert rune slices into strings
	gammaRateBinary := string(gammaRateBits)
	epsilonRateBinary := string(epsilonRateBits)

	// Convert binary strings into decimal unsigned integers
	gammaRate, _ := strconv.ParseUint(gammaRateBinary, 2, 64)
	epsilonRate, _ := strconv.ParseUint(epsilonRateBinary, 2, 64)

	powerConsumption := gammaRate * epsilonRate

	fmt.Println(powerConsumption)
}

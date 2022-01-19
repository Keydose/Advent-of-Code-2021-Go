package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// TODO: Load input from txt file, loop through each line, count the amount of times it's greater than the prev. measurement
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

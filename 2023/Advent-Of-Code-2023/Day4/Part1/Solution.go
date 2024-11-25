package main

import (
	"fmt"
	"os"
	"strings"
)

func readAndExtract(filename string) ([]string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")

	var results []string
	for _, line := range lines {
		parts := strings.Split(line, ":")
		if len(parts) > 1 {
			results = append(results, strings.TrimSpace(parts[1]))
		}
	}

	return results, nil
}

func pow(n int, exp int) int {
	if exp < 0 { // handle negative exponents
		return 0
	}
	result := 1
	for i := 0; i < exp; i++ {
		result *= n
	}
	return result
}

func checkWinningNumbers(round string) int {
	cards := strings.Split(round, "|")
	winningNumbers := strings.Fields(strings.TrimSpace(cards[0]))
	ourNumbers := strings.Fields(strings.TrimSpace(cards[1]))

	// fmt.Printf("[Winning]: %#q", winningNumbers)
	// fmt.Printf("[Our]: %#q", ourNumbers)

	var pts int
	for _, oNum := range ourNumbers {
		for _, wNum := range winningNumbers {
			if oNum == wNum {
				pts++
				break
			}
		}
	}

	return pow(2, pts-1)
}

func main() {
	filename := "input.txt"
	// filename := "test.txt"
	var total int

	rounds, err := readAndExtract(filename)
	for _, round := range rounds {
		total += checkWinningNumbers(round)
	}

	if err != nil {
		fmt.Println("Error")
		return
	}

	fmt.Printf("Answer: [%v]\n", total) // Answer: 21558
}

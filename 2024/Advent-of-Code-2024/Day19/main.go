package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	content, _ := os.ReadFile("input.in")
	fileContent := strings.Split(string(content), "\r\n\r\n")

	patterns := strings.Split(fileContent[0], ", ")
	designs := strings.Fields(fileContent[1])

	solveA(patterns, designs)
	// solveB(patterns, designs)

	elapsed := time.Since(start)
	fmt.Println("Execution time: ", elapsed)
}

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
	fileContent := strings.Fields(string(content))
	grid := make([][]rune, len(fileContent))
	for i, line := range fileContent {
		grid[i] = []rune(line)
	}
	solveA(grid)
	// solveB(fileContent)
	elapsed := time.Since(start)
	fmt.Println("Execution time: ", elapsed)
}

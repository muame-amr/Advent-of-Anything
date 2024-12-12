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
	// solveA(fileContent)
	solveB(fileContent)
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s\n", elapsed)
}

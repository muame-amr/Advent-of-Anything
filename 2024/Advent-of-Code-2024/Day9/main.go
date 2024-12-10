package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	content, _ := os.ReadFile("test.in")
	fileContent := strings.TrimSpace(string(content))
	// solveA(fileContent)
	solveB(fileContent)
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s\n", elapsed)
}

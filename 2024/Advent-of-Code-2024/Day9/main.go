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
	fileContent := strings.TrimSpace(string(content))
	solveA(fileContent)
	// solveB(fileContent)
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s\n", elapsed)
}

// result:
// 6370402949053
// 6370402949053
// test: 523702

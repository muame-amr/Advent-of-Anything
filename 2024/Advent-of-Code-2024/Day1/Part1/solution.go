package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var first, second []float64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())

		st, err1 := strconv.ParseFloat(parts[0], 64)
		nd, err2 := strconv.ParseFloat(parts[1], 64)

		if err1 == nil && err2 == nil {
			first = append(first, st)
			second = append(second, nd)
		}
	}

	// Check for any scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	sort.Float64s(first)
	sort.Float64s(second)

	total := 0
	for i := 0; i < len(first); i++ {
		total += int(math.Abs(first[i] - second[i]))
	}

	fmt.Println(total)
}

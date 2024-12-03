package main

import (
	"bufio"
	"fmt"
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

	var first, second []int
	ctr := make(map[int]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())

		st, err1 := strconv.Atoi(parts[0])
		nd, err2 := strconv.Atoi(parts[1])

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

	sort.Ints(first)
	sort.Ints(second)

	for i := 0; i < len(second); i++ {
		ctr[second[i]] += 1
	}

	total := 0
	for i := 0; i < len(first); i++ {
		total += first[i] * ctr[first[i]]
	}

	fmt.Println(total)
}

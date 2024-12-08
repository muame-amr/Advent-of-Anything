package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	content, _ := os.ReadFile("input.in")
	fileContent := strings.Split(string(content), "\r\n")
	solveA(fileContent)
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s\n", elapsed)
}

func solveA(fileContent []string) {
	var eval func(test int, nums []string, idx int, res int) bool
	eval = func(test int, nums []string, idx int, res int) bool {
		if idx == len(nums) {
			return res == test
		}

		num, _ := strconv.Atoi(nums[idx])

		return eval(test, nums, idx+1, res+num) || eval(test, nums, idx+1, res*num)
	}

	total := 0

	for _, line := range fileContent {
		test, _ := strconv.Atoi(strings.Split(line, ":")[0])
		nums := strings.Fields(strings.Split(line, ":")[1])

		if eval(test, nums, 0, 0) {
			fmt.Println("Success test: ", test)
			total += test
		}
	}

	fmt.Println(total)
}

func solveB(fileContent []string) {
	total := 0
	for _, line := range fileContent {
		test, _ := strconv.Atoi(strings.Split(line, ":")[0])
		nums := strings.Fields(strings.Split(line, ":")[1])

		opCombinations := generateOperatorCombinations(len(nums))

		for _, combination := range opCombinations {
			res, _ := strconv.Atoi(nums[0])
			for i, op := range combination {
				nextNum, _ := strconv.Atoi(nums[i+1])
				if op == "+" {
					res += nextNum
				} else if op == "*" {
					res *= nextNum
				}
			}
			if res == test {
				total += test
				break
			}
		}
	}
	fmt.Println(total)
}

func generateOperatorCombinations(n_size int) [][]string {
	operators := []string{"+", "*"}
	var combinations [][]string

	operatorCount := n_size - 1
	totalCombinations := intPow(len(operators), operatorCount)

	for i := 0; i < totalCombinations; i++ {
		current := make([]string, operatorCount)
		for j := 0; j < operatorCount; j++ {
			current[j] = operators[(i>>j)%len(operators)]
		}
		combinations = append(combinations, current)
	}
	return combinations
}

func intPow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}
